package main

import (
	"log"
	"net/http"
)

const port = ":4000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Printf("Starting server on %s", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
