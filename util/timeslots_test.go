package util

import "testing"

var tsm = NewTimeslotMap()

func TestGet(t *testing.T) {
	slotnumber := tsm.Get("1", "9", "10")
	if slotnumber != "1" {
		t.Errorf("Slot number was incorrect, got: %s, want: %s.", slotnumber, "1")
	}
}

func TestCourt5Time1945(t *testing.T) {
	slotnumber := tsm.Get("5", "19", "45")
	if slotnumber != "97" {
		t.Errorf("Slot number was incorrect, got: %s, want: %s.", slotnumber, "97")
	}
}
