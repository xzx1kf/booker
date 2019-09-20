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
	"golang.org/x/net/publicsuffix"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	tynemouthSquashUrl = "http://tynemouth-squash.herokuapp.com/bookings"
)

type Court struct {
	Court    string `json:"court"`
	Days     string `json:"days"`
	Hour     string `json:"hour"`
	Min      string `json:"min"`
	Timeslot string `json:"timeslot"`
}

func HandleRequest() error {

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
		return errors.New("failed to recieve messages from queue")
	}

	if len(result.Messages) == 0 {
		return errors.New("no messages to process")
	}

	msg := result.Messages[0]
	event := Court{}
	json.Unmarshal([]byte(*msg.Body), &event)

	// Delete the message from the queue.
	_, err = svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &qURL,
		ReceiptHandle: result.Messages[0].ReceiptHandle,
	})

	if err != nil {
		return errors.New("Failed to delete the message from queue")
	}

	fmt.Println("Message Deleted")

	// create a cookiejar - this is required because the website uses cookies
	// and without it the booking of a court fails
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return err
	}

	client := &http.Client{
		Jar: jar,
	}

	// Get the court booking page - this creates the cookie
	r, err := http.NewRequest("GET", tynemouthSquashUrl+"/new?"+
		"court="+event.Court+
		"&days="+event.Days+
		"&hour="+event.Hour+
		"&min="+event.Min+
		"&timeSlot="+event.Timeslot,
		nil)
	if err != nil {
		return err
	}

	// Make the request
	rsp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	token, time := parseCourtBookingPage(doc)

	v := url.Values{}
	v.Set("utf8", "&#x2713;")
	v.Set("authenticity_token", token)
	v.Set("booking[full_name]", os.Getenv("NAME"))
	v.Set("booking[membership_number]", os.Getenv("MEMBERSHIP_NUMBER"))
	v.Set("booking[vs_player_name]", "")
	v.Set("booking[booking_number]", "1")
	v.Set("booking[start_time]", time)
	v.Set("booking[time_slot_id]", event.Timeslot)
	v.Set("booking[court_time]", "40")
	v.Set("booking[court_id]", event.Court)
	v.Set("booking[days]", event.Days)
	v.Set("commit", "Book Court")

	// Create the POST request.
	r, err = http.NewRequest("POST", tynemouthSquashUrl, strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	// Perform the POST request
	rsp, err = client.Do(r)
	if err != nil {
		return err
	}

	u, err := url.Parse(rsp.Request.URL.String())
	if err != nil {
		return err
	}

	// if the response url contains an error parameter then the booking
	// must of failed.
	m, _ := url.ParseQuery(u.RawQuery)
	if m.Get("error") != "" {
		err := fmt.Errorf("Court %s is already booked at %s", event.Court, time)
		return err
	}
	fmt.Println("Court", event.Court, "booked at", time)
	return nil
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
	lambda.Start(HandleRequest)
}
