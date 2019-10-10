package tynemouth

import (
	"fmt"
	"net/url"
	"testing"
)

func TestListSlot(t *testing.T) {
	client := NewClient(nil)
	url, _ := url.Parse(defaultBaseURL)
	client.BaseURL = url

	slot := &Slot{Court: "1", Days: "20", Hour: "19", Min: "50", TimeSlot: "17"}
	authToken := "token"
	time := "2019-10-30 19:50:00 +0000"

	err := client.Slot.ListSlot(slot)
	if err != nil {
		fmt.Println(err)
	}

	// test that relative URL was expanded
	if got, want := slot.AuthToken, authToken; got != want {
		t.Errorf("ListSlot() AuthToken is %v, want %v", got, want)
	}

	// test that time was captured
	if got, want := slot.Time, time; got != want {
		t.Errorf("ListSlot() Time is %v, want  %v", got, want)
	}
}
