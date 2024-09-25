package main

import (
	"fmt"
	"net/http"
)

type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 not found: %s\n", r.URL)
	}
}

func main() {
	e := &Engine{}
	http.ListenAndServe(":8080", e)
}
