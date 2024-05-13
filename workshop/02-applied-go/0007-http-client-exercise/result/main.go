package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	const api = "https://themealdb.com/api/json/v1/1/random.php"
	resp, err := http.Get(api)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	fmt.Println("Response status:", resp.Status)

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response body:", string(body))
}
