package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
}
