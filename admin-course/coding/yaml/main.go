package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {
	yfile, err := ioutil.ReadFile("items.yaml")

	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]int)

	err = yaml.Unmarshal(yfile, &data)

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range data {
		fmt.Printf("%s -> %d\n", k, v)
	}
}
