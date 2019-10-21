package main

import (
	"testing"
)

func TestPathDay(t *testing.T) {
	p := NewPath("/courts/1")

	// test that relative URL was expanded
	if got, want := p.ID, "1"; got != want {
		t.Errorf("NewPath() ID is %v, want %v", got, want)
	}
}

func TestPathId(t *testing.T) {
	p := NewPath("/courts/1/12")

	// test that relative URL was expanded
	if got, want := p.Slot, "12"; got != want {
		t.Errorf("NewPath() Slot is %v, want %v", got, want)
	}
}
