package main

import (
	"errors"
	"net/http"
)

type slot struct {
	ID      string `json:"id"`
	AuthKey string `json:"authkey"`
}

func (s *Server) handleSlots(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleSlotsGet(w, r)
		return
	case "POST":
		s.handleSlotsPost(w, r)
		return
	}
	respondHTTPErr(w, r, http.StatusNotFound)
}

func (s *Server) handleSlotsGet(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError,
		errors.New("not implemented"))
}

func (s *Server) handleSlotsPost(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError,
		errors.New("not implemented"))
}
