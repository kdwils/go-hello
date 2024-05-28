package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)
	// http.HandleFunc("/canary", RolloutHandler)
	http.ListenAndServe(":8080", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func RolloutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from the preview version, %s!", r.URL.Path[1:])
}
