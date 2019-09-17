package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/publicsuffix"
)

const (
	tynemouthSquashUrl = "http://tynemouth-squash.herokuapp.com/bookings"
)

type Booking struct {
	Court       string `json:"court"`
	Days        string `json:"days"`
	Hour        string `json:"hour"`
	Min         string `json:"min"`
	Timeslot    string `json:"hour"`
	PlayerA     string `json:"playerA"`
	PlayerB     string `json:"playerB"`
	BookingLink string `json:"bookingLink"`
	Booked      bool   `json:"booked"`
}

type Bookings []Booking

func bookCourt(court, days, hour, min, timeslot string) {
	// Book a court given the
	//  - Court Number
	//  - Time
	//  - Date
	//  - etc...

	// create a cookiejar - this is required because the website uses cookies
	// and without it the booking of a court fails
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	// Make the request
	rsp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	// Perform the POST request
	rsp, err = client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

}

func listAvailableCourts() {
	// Rename to List Bookings
	bookings := Bookings{}

	res, err := http.Get("http://tynemouth-squash.herokuapp.com/bookings?day=21")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Available Courts")
	doc.Find(".booking div.book a.booking_link").Each(func(i int, s *goquery.Selection) {
		bl, exists := s.Attr("href")
		if exists {
			bs := parseBookingUrl(bl)

			bookings = append(bookings, bs)

            // book court on saturdays
			if bs.Court == "1" && bs.Days == "21" && bs.Hour == "9" && bs.Min == "10" && bs.Timeslot == "1" {
				bookCourt(bs.Court, bs.Days, bs.Hour, bs.Min, bs.Timeslot)
				fmt.Println("booked")
			}
		}
	})
}

func parseBookingUrl(link string) Booking {
	s := link[14:]
	s = strings.Replace(s, "&amp", "", -1)
	m, err := url.ParseQuery(s)
	if err != nil {
		log.Fatal(err)
	}

	bs := Booking{
		Court:       m["court"][0],
		Days:        m["days"][0],
		Hour:        m["hour"][0],
		Min:         m["min"][0],
		Timeslot:    m["timeSlot"][0],
		Booked:      false,
		BookingLink: "http://tynemouth-squash.herokuapp.com" + link,
	}

	return bs
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
	listAvailableCourts()
}
