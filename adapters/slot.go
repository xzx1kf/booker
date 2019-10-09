package tynemouth

import (
	"fmt"
	"net/http"
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

func (s *Slot) Parse(resp *http.Response) {
	s.AuthToken = "token"
	s.Time = "19:50"
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

	// slot needs authtoken and time populating in the write method
	_, err = s.client.Do(req, opt)
	if err != nil {
		return err
	}

	return nil

}
