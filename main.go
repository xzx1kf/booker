package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/publicsuffix"
)

const (
	tynemouthSquashUrl = "http://tynemouth-squash.herokuapp.com/bookings"
)

func bookCourt(court, days, hour, min, timeslot string) (message string, err error) {

	// create a cookiejar - this is required because the website uses cookies
	// and without it the booking of a court fails
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return "Failed to book court.", err
	}

	client := &http.Client{
		Jar: jar,
	}

	// Get the court booking page - this creates the cookie
	r, err := http.NewRequest("GET", tynemouthSquashUrl+"/new?"+
		"court="+court+
		"&days="+days+
		"&hour="+hour+
		"&min="+min+
		"&timeSlot="+timeslot,
		nil)
	if err != nil {
		return "Failed to book court.", err
	}

	// Make the request
	rsp, err := client.Do(r)
	if err != nil {
		return "Failed to book court.", err
	}
	defer rsp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(rsp.Body)
	token, time := parseCourtBookingPage(doc)

	v := url.Values{}
	v.Set("utf8", "&#x2713;")
	v.Set("authenticity_token", token)
	v.Set("booking[full_name]", "Nick Hale")
	v.Set("booking[membership_number]", "s119")
	v.Set("booking[vs_player_name]", "")
	v.Set("booking[booking_number]", "1")
	v.Set("booking[start_time]", time)
	v.Set("booking[time_slot_id]", timeslot)
	v.Set("booking[court_time]", "40")
	v.Set("booking[court_id]", court)
	v.Set("booking[days]", days)
	v.Set("commit", "Book Court")

	// Create the POST request.
	r, err = http.NewRequest("POST", tynemouthSquashUrl, strings.NewReader(v.Encode()))
	if err != nil {
		return "Failed to book court.", err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	// Perform the POST request
	rsp, err = client.Do(r)
	if err != nil {
		return "Failed to book court.", err
	}

	u, err := url.Parse(rsp.Request.URL.String())
	if err != nil {
		return "Failed to book court.", err
	}

	// if the response url contains an error parameter then the booking
	// must of failed.
	m, _ := url.ParseQuery(u.RawQuery)
	if m.Get("error") != "" {
		return "Failed to book court. The court already has a booking at this time", errors.New("Court already booked")
	}
	return "Court booked.", nil
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
	courtPtr := flag.String("c", "", "Which court to book.")
	daysPtr := flag.String("d", "", "Which day to book on d=0 being today.")
	hourPtr := flag.String("h", "", "What time, the hour portion in 24hr format 0-23.")
	minPtr := flag.String("m", "", "What time, the minute portion from 0-59.")
	tsPtr := flag.String("t", "", "What timeslot is the court.")
	flag.Parse()

	if *courtPtr == "" || *hourPtr == "" || *minPtr == "" || *daysPtr == "" || *tsPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("court: %s, time: %s:%s, days: %s, timeslot: %s\n", *courtPtr, *hourPtr, *minPtr, *daysPtr, *tsPtr)

	message, _ := bookCourt(*courtPtr, *daysPtr, *hourPtr, *minPtr, *tsPtr)
	fmt.Println(message)
}
