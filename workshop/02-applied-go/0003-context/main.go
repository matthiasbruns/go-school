package main

import (
	"context"
	"fmt"
	"net/http"
)

func doSomething(ctx context.Context) error {
	// check if the context has a value
	if v := ctx.Value("greeted"); v != nil {
		fmt.Println("greeted:", v)
		return nil
	}

	return fmt.Errorf("not greeted")
}

func hello(w http.ResponseWriter, r *http.Request) {
	// in requests, you get the context from the request
	// it also handles timeouts and cancellations
	ctx := r.Context()

	ctx = context.WithValue(ctx, "greeted", true)

	// send the context to the function
	if err := doSomething(ctx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello,")
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
