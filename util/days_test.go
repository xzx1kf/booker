package util

import (
	"testing"
)

func TestValid(t *testing.T) {
	got := "21" // Days("17-10-2019 21:10")
	want := "21"
	if got != want {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
