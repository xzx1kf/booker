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

	fmt.Println(client.BaseURL)

	slot := Slot{Court: "1", Days: "21", Hour: "19", Min: "50", TimeSlot: "17"}

	err := client.Slot.ListSlot(slot)
	if err != nil {
		fmt.Println(err)
	}

}
