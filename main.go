package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/xzx1kf/booker/timeslots"
	"golang.org/x/net/publicsuffix"
)

const (
	tynemouthSquashUrl = "http://tynemouth-squash.herokuapp.com/bookings"
)

type Court struct {
	Court string `json:"court"`
	Days  string `json:"days"`
	Hour  string `json:"hour"`
	Min   string `json:"min"`
}

func BookCourts() error {
	// read the booking info from SQS i.e. Court, Days, Hour, Min
	booking, err := getBookingInfo()

	// lookup the timeslot value given the Court, Hour and Min to be booked
	timeslots.Init() // create the lookup map
	timeslot := timeslots.Get(booking.Court, booking.Hour, booking.Min)

	// create a cookiejar
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return err
	}

	client := &http.Client{
		Jar: jar,
	}

	doc, err := getBookingPage(client, booking, timeslot)
	if err != nil {
		return err
	}

	token, time := parseCourtBookingPage(doc)

	v := url.Values{}
	v.Set("utf8", "&#x2713;")
	v.Set("authenticity_token", token)
	v.Set("booking[full_name]", os.Getenv("NAME"))
	v.Set("booking[membership_number]", os.Getenv("MEMBERSHIP_NUMBER"))
	v.Set("booking[vs_player_name]", "")
	v.Set("booking[booking_number]", "1")
	v.Set("booking[start_time]", time)
	v.Set("booking[time_slot_id]", timeslot)
	v.Set("booking[court_time]", "40")
	v.Set("booking[court_id]", booking.Court)
	v.Set("booking[days]", booking.Days)
	v.Set("commit", "Book Court")

	// Create the POST request.
	request, err := http.NewRequest("POST", tynemouthSquashUrl, strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	// Perform the POST request
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	u, err := url.Parse(response.Request.URL.String())
	if err != nil {
		return err
	}

	// if the response url contains an error parameter then the booking
	// must of failed.
	m, _ := url.ParseQuery(u.RawQuery)
	if m.Get("error") != "" {
		err := fmt.Errorf("Court %s is already booked at %s", booking.Court, time)
		return err
	}
	fmt.Println("Court", booking.Court, "booked at", time)
	return nil
}

func getBookingPage(client *http.Client, booking Court, timeslot string) (doc *goquery.Document, err error) {
	request, err := http.NewRequest("GET",
		tynemouthSquashUrl+"/new?"+
			"court="+booking.Court+
			"&days="+booking.Days+
			"&hour="+booking.Hour+
			"&min="+booking.Min+
			"&timeSlot="+timeslot,
		nil)
	if err != nil {
		return nil, err
	}

	// Make the request
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	doc, err = goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func getBookingInfo() (Court, error) {
	svc := sqs.New(session.New())
	qURL := os.Getenv("BOOKING_QUEUE")
	event := Court{}

	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &qURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(20), // 20 seconds
		WaitTimeSeconds:     aws.Int64(0),
	})
	if err != nil {
		return event, errors.New("failed to recieve messages from queue")
	}

	if len(result.Messages) == 0 {
		return event, errors.New("no messages to process")
	}

	msg := result.Messages[0]

	json.Unmarshal([]byte(*msg.Body), &event)

	// Delete message from SQS
	_, err = svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &qURL,
		ReceiptHandle: result.Messages[0].ReceiptHandle,
	})
	if err != nil {
		return event, errors.New("Failed to delete the message from queue")
	}

	fmt.Println("Message Deleted")

	return event, nil
}

func parseCourtBookingPage(doc *goquery.Document) (token string, time string) {
	s := doc.Find("form.booking")
	s.Find("input").Each(func(i int, sel *goquery.Selection) {
		input, _ := sel.Attr("name")
		if input == "authenticity_token" {
			token, _ = sel.Attr("value")
		} else if input == "booking[start_time]" {
			time, _ = sel.Attr("value")
		}
	})

	return token, time
}

func main() {
	lambda.Start(BookCourts)
}
