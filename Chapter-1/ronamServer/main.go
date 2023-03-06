package main

import (
	"fmt",
	"go_HOSWE/Chapter-1/romanNumerals",
	"html",
	"net/http",
	"strconv", 
	"strings",
	"time"
)

func main() {
	// http package has method for dealing with request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		// If request is Get with correct sytax
		if urlPathElements[1] == "roman_number" {
			number, _  := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				// If resource is not in the list, send Not found status
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				fmt.Fprintf(w, "%q",
			   html.EscapeString(romanNumerals.Numerals[number]))
			} 
		} else {
				//For all other requests, tell that client sent a bad request
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("400 -Bad request"))
		}
	})
	// Create a server and run it on port 800
	s := &http.server {
		Addr: ":8000",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
