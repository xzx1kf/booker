package tynemouth

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

type SlotService service

type Parser interface {
	Parse(resp *http.Response)
}

type Slot struct {
	Court     string `url:"court"`
	Days      string `url:"days"`
	Hour      string `url:"hour"`
	Min       string `url:"min"`
	TimeSlot  string `url:"timeSlot"`
	AuthToken string `url:"auth,omitempty"`
	Time      string `url:"time,omitempty"`
}

// Parse the tynemouth squash new booking page. We need to extract
// the authenticity_token and booking[start_time]
// TODO investigate whether booking[start_time] could be a given.
func (s *Slot) Parse(resp *http.Response) {
	tokenizer := html.NewTokenizer(resp.Body)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
		}

		if tokenType == html.SelfClosingTagToken {
			token := tokenizer.Token()
			if "input" == token.Data {
				ok, authToken := getAttribute(token, "authenticity_token")
				if ok {
					s.AuthToken = authToken
					continue
				}

				ok, time := getAttribute(token, "booking[start_time]")
				if ok {
					s.Time = time
					continue
				}
			}
		}
	}
}

func getAttribute(t html.Token, attrName string) (ok bool, value string) {
	found := false

	for _, v := range t.Attr {
		if v.Key == "name" && v.Val == attrName {
			found = true
			break
		}
	}

	if found {
		for _, v := range t.Attr {
			if v.Key == "value" {
				return true, v.Val
			}
		}
	}

	return
}

func (s *SlotService) ListSlot(opt *Slot) error {
	u := fmt.Sprintf("%v%v/%v", defaultBaseURL, "bookings", "new")
	u, err := addOptions(u, opt)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest("GET", u)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, opt)
	if err != nil {
		return err
	}

	return nil

}
