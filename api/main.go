package main

import (
	"context"
	"flag"
	"log"
	"net/http"
)

type Server struct {
}

type contextKey struct {
	name string
}

var contextKeyAuthKey = &contextKey{"pGbB+Um0cMCrZYVfCTVOFbpu1NyMjIcz3b7+R8xgFrrC9435ojNrfv5RKIHtwEYA6/sXHHU/GUBrpEOLVMeICQ=="}

func main() {
	var (
		addr = flag.String("addr", ":8080", "endpoint address")
	)
	s := &Server{}
	mux := http.NewServeMux()
	mux.HandleFunc("/courts/", withCORS(withAuthKey(s.handleSlots)))
	log.Println("Starting web server on", *addr)
	http.ListenAndServe(":8080", mux)
	log.Println("Stopping...")
}

func AuthKey(ctx context.Context) (string, bool) {
	key, ok := ctx.Value(contextKeyAuthKey).(string)
	return key, ok
}

func withAuthKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if !isValidAuthKey(key) {
			respondErr(w, r, http.StatusUnauthorized, "invalid auth key")
			return
		}
		ctx := context.WithValue(r.Context(), contextKeyAuthKey, key)
		fn(w, r.WithContext(ctx))
	}
}

func isValidAuthKey(key string) bool {
	return key == "pGbB+Um0cMCrZYVfCTVOFbpu1NyMjIcz3b7+R8xgFrrC9435ojNrfv5RKIHtwEYA6/sXHHU/GUBrpEOLVMeICQ=="
}

func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Location")
		fn(w, r)
	}
}
