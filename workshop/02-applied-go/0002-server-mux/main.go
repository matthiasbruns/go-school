package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello,")
}

func world(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/world", world)

	if err := http.ListenAndServe(":8090", mux); err != nil {
		panic(err)
	}
}
