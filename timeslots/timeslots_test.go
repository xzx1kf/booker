package timeslots

import "testing"

func TestGet(t *testing.T) {
    Init()
    slotnumber := Get("1", "9", "10")
    if slotnumber != "1" {
        t.Errorf("Slot number was incorrect, got: %s, want: %s.", slotnumber, "1")
    }
}
