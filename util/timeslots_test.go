package util

import "testing"

/*
var tsm = NewTimeslotMap()

func TestGet(t *testing.T) {
	slotnumber := tsm.Lookup("1", "9", "10")
	if slotnumber != "1" {
		t.Errorf("Slot number was incorrect, got: %s, want: %s.", slotnumber, "1")
	}
}

func TestCourt5Time1945(t *testing.T) {
	slotnumber := tsm.Lookup("5", "19", "45")
	if slotnumber != "97" {
		t.Errorf("Slot number was incorrect, got: %s, want: %s.", slotnumber, "97")
	}
}

func TestInvalidLookup(t *testing.T) {
	slotnumber := tsm.Lookup("6", "9", "10")
	if slotnumber != "" {
		t.Errorf("Slot number was incorrect, got: %s, want %s.", slotnumber, "")
	}
}
*/
func TestTimeslot(t *testing.T) {
	slotnumber := Timeslot("1", "10/10/2019 09:50")
	if slotnumber != "2" {
		t.Errorf("Slot number was incorrect, got: %s, want %s.", slotnumber, "2")
	}
}

func TestTimeslotCourt2(t *testing.T) {
	slotnumber := Timeslot("2", "10/10/2019 09:50")
	want := "22"
	if slotnumber != want {
		t.Errorf("Slot number was incorrect, got: %s, want %s.", slotnumber, want)
	}
}

func TestTimeslotCourt3(t *testing.T) {
	slotnumber := Timeslot("3", "10/10/2019 10:20")
	want := "43"
	if slotnumber != want {
		t.Errorf("Slot number was incorrect, got: %s, want %s.", slotnumber, want)
	}
}
