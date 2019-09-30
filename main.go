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
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/xzx1kf/booker/util"
	"golang.org/x/net/publicsuffix"
)

const (
	tynemouthSquashUrl = "http://tynemouth-squash.herokuapp.com/bookings"
)

type Court struct {
	Id   string `json:"id"`
	Date string `json:"date"`
}

func BookCourts() (err error) {
	// Read the booking info from SQS.
	court, err := getBookingInfo()
	if err != nil {
		return err
	}

	// Parse the date string into date, time
	layout := "02/01/2006 15:04"
	date, _ := time.Parse(layout, court.Date)
	hour := date.Format("15")
	min := date.Format("04")

	// TODO: use the days value to determine whether the court can be booked
	// today. If days >= 22 don't book and don't delete from the queue.
	// Deletion of the message from the queue is done in getBookingInfo()
	days := util.Days(date)

	// Lookup the timeslot value.
	timeslot := util.Timeslot(court.Id, court.Date)

	// Initialize a http.Client and add a cookiejar.
	// For some reason the POST method will fail if this cookie jar
	// is not present.
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	client := &http.Client{Jar: jar}

	// Make the http 'GET' request to get the court booking page.
	doc, err := getBookingPage(client, court.Id, days, hour, min, timeslot)
	if err != nil {
		return err
	}

	// Scrape the time and authenticity token from the booking page.
	token, time := parseCourtBookingPage(doc)

	// Make the http 'POST' request to book the court.
	err = postRequest(client, court.Id, days, token, time, timeslot)
	if err != nil {
		return err
	}

	fmt.Println("Court:", court.Id, "booked at:", court.Date)
	return
}

func postRequest(client *http.Client, id, days, token, time, timeslot string) error {
	data := url.Values{}
	data.Set("utf8", "&#x2713;")
	data.Set("authenticity_token", token)
	data.Set("booking[full_name]", os.Getenv("NAME"))
	data.Set("booking[membership_number]", os.Getenv("MEMBERSHIP_NUMBER"))
	data.Set("booking[vs_player_name]", "")
	data.Set("booking[booking_number]", "1")
	data.Set("booking[start_time]", time)
	data.Set("booking[time_slot_id]", timeslot)
	data.Set("booking[court_time]", "40")
	data.Set("booking[court_id]", id)
	data.Set("booking[days]", days)
	data.Set("commit", "Book Court")

	// Create the POST request.
	request, _ := http.NewRequest("POST", tynemouthSquashUrl, strings.NewReader(data.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	// Inspect the url in the response to see if it contains an error parameter.
	// This can be used to determine if the court booking was successful.
	u, err := url.Parse(response.Request.URL.String())
	if err != nil {
		return err
	}

	// if the response url contains an error parameter then the booking
	// must of failed.
	m, _ := url.ParseQuery(u.RawQuery)
	if m.Get("error") != "" {
		err := fmt.Errorf("Court %s is already booked at %s", id, time)
		return err
	}
	return nil
}

func getBookingPage(client *http.Client, id, days, hour, min, timeslot string) (doc *goquery.Document, err error) {
	request, err := http.NewRequest("GET",
		tynemouthSquashUrl+"/new?"+
			"court="+id+
			"&days="+days+
			"&hour="+hour+
			"&min="+min+
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

func getBookingInfo() (event Court, err error) {
	svc := sqs.New(session.New())
	qURL := os.Getenv("BOOKING_QUEUE")

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
