package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize servemux (router)
	mux := http.NewServeMux()

	// Create file server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Log statements
	log.Println("starting server on :4000")

	// Start a new web server
	// Parameters: TCP network adress to listen,and servemux
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
