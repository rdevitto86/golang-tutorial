package main

import (
	"fmt"
	"log"
	"net/http"
)

// [1:] creates a "sub-slice" of the URL path. string starts at index 1 to avoid "/" char
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// build and run the program and go to http://localhost:8080/[insert string here]
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
