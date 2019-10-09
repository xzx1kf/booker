package tynemouth

import (
	"fmt"
)

type SlotService service

type Slot struct {
	Court     string `url:"court"`
	Days      string `url:"days"`
	Hour      string `url:"hour"`
	Min       string `url:"min"`
	TimeSlot  string `url:"timeSlot"`
	AuthToken string `url:"auth,omitempty"`
	Time      string `url:"time,omitempty"`
}

func (s Slot) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (s *SlotService) ListSlot(opt Slot) error {
	u := fmt.Sprintf("%v%v/%v", defaultBaseURL, "bookings", "new")
	u, err := addOptions(u, opt)
	if err != nil {
		return err
	}

	fmt.Println(u)

	req, err := s.client.NewRequest("GET", u)
	if err != nil {
		return err
	}

	// Slot needs to implement the io.Writer interface
	var slot *Slot
	resp, err := s.client.Do(req, &slot)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	fmt.Println("HELLO")
	return nil

}
