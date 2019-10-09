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

	err := client.Slot.ListSlot(slot)
	if err != nil {
		fmt.Println(err)
	}

	// test that relative URL was expanded
	if got, want := slot.AuthToken, authToken; got != want {
		t.Errorf("ListSlot() AuthToken is %v, want %v", got, want)
	}
}
