package tynemouth

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type Slot struct {
	Court string `json:"court"`
	Date  string `json:"date"`
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil)

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	inBody, outBody := Slot{Court: "1", Date: "10/10/2019"}, `{"court":"1","date":"10/10/2019"}`+"\n"

	req, _ := c.NewRequest("GET", inURL, inBody)

	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}
	got, want := req.URL.String(), outURL
	fmt.Println(got, want)

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}
}
