package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	//START_1 OMIT
	resp, err := http.Get("http://google.com") // HL1
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	// bufio.NewScanner reads each row of the body
	scanner := bufio.NewScanner(resp.Body) // HL1
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text()) // prints each row of the body
	}
	//END_1 OMIT

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
