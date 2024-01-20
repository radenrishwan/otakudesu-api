package utils

import (
	"log"
	"net/http"
)

// ErrorHandler is handler for handle error
func LoggerHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)

		h.ServeHTTP(w, r)
	})
}
