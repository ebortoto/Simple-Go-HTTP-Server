package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Hello request from %s\n", req.RemoteAddr)
	fmt.Fprintf(w, "hello\n")
}

func wait(w http.ResponseWriter, req *http.Request) {
	// 1. Extract the text after "/wait/"
	// If URL is "/wait/5", text will be "5"
	text := strings.TrimPrefix(req.URL.Path, "/wait/")

	// 2. Convert text to a number (integer)
	seconds, err := strconv.Atoi(text)
	if err != nil || seconds < 0 {
		http.Error(w, "Invalid number of seconds", http.StatusBadRequest)
		return
	}

	fmt.Printf("Sleeping for %d seconds...\n", seconds)

	// 3. Sleep for the requested duration
	time.Sleep(time.Duration(seconds) * time.Second)

	// 4. Respond
	fmt.Fprintf(w, "waited %d seconds\n", seconds)
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/wait/", wait)

	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
