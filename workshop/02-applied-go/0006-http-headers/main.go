package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8090", nil)

	req.Header.Add("X-Request-ID", "123")
	req.SetBasicAuth("user", "password")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println(resp.Status)
}
