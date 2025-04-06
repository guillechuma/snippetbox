package main

import (
	"log"
	"net/http"
)

// Home handler function
// Parameters:
// - http.ResponseWriter: HTTP response
// - *http.Request: pointer to the current request
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Initialize servemux (router)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Log statements
	log.Println("starting server on :4000")

	// Start a new web server
	// Parameters: TCP network adress to listen,and servemux
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
