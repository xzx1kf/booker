package tynemouth

import (
	"fmt"
	"net/http"
	"testing"
)

type Slot struct {
	Court string `json:"court"`
	Date  string `json:"date"`
}

type BookingInfo struct {
	Court    string `url:"court"`
	Days     string `url:"days"`
	Hour     string `url:"hour"`
	Min      string `url:"min"`
	TimeSlot string `url:"timeSlot"`
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	inURL, outURL := "/foo", defaultBaseURL+"foo"

	req, _ := c.NewRequest(http.MethodGet, inURL)

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that default user-agent is attaached to the request
	if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
		t.Errorf("NewRequest() User-Agent is %v, want %v", got, want)
	}
}

func TestBuildUrl(t *testing.T) {
	info := BookingInfo{Court: "1", Days: "21", Hour: "19", Min: "50", TimeSlot: "17"}

	u := fmt.Sprintf("%v%v/%v", defaultBaseURL, "bookings", "new")
	s, _ := addOptions(u, info)

	outURL := defaultBaseURL + "bookings/new?court=1&days=21&hour=19&min=50&timeSlot=17"

	if got, want := s, outURL; got != want {
		t.Errorf("AddOptions() URL is %v, want %v", got, want)
	}
}
