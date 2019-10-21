package main

import (
	"strings"
)

const PathSeparator = "/"

type Path struct {
	Path string
	ID   string
	Slot string
}

func NewPath(p string) *Path {
	var id string
	var slot string
	p = strings.Trim(p, PathSeparator)
	s := strings.Split(p, PathSeparator)
	if len(s) > 2 {
		id = s[len(s)-2]
		slot = s[len(s)-1]
		p = strings.Join(s[:len(s)-2], PathSeparator)
	} else if len(s) > 1 {
		id = s[len(s)-1]
		p = strings.Join(s[:len(s)-1], PathSeparator)
	}
	return &Path{Path: p, ID: id, Slot: slot}
}

func (p *Path) HasID() bool {
	return len(p.ID) > 0
}
