package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

//START_1 OMIT
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n") // HL1
}

func main() {
	http.HandleFunc("/hello", hello) // HL1
	http.HandleFunc("/headers", headers)

	fmt.Println("Server is up - http://localhost:8090")
	http.ListenAndServe(":8090", nil) // HL1
}

//END_1 OMIT
