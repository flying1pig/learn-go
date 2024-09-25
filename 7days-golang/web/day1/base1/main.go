package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexH)
	http.ListenAndServe(":8080", nil)
}

func indexH(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
