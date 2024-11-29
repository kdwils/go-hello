package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/stable/", HelloHandler)
	// http.HandleFunc("/rollout/", RolloutHandler)
	http.ListenAndServe(":8080", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
