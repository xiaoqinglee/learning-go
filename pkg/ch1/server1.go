package ch1

import (
	"fmt"
	"log"
	"net/http"
)

func Server1() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q", r.URL.Path)
}
