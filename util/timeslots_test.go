package util

import "testing"

func TestGet(t *testing.T) {
	Init()
	slotnumber := Get("1", "9", "10")
	if slotnumber != "1" {
		t.Errorf("Slot number was incorrect, got: %s, want: %s.", slotnumber, "1")
	}
}

func TestCourt5Time1945(t *testing.T) {
	Init()
	slotnumber := Get("5", "19", "45")
	if slotnumber != "97" {
		t.Errorf("Slot number was incorrect, got: %s, want: %s.", slotnumber, "97")
	}
}
