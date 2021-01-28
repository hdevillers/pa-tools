package main

import (
	"fmt"
	"log"

	"github.com/hdevillers/pa-tools/sam"
)

func main() {
	log.SetPrefix("Parsing: ")
	value, err := sam.CigarConsume("150M")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %d\n", value)
}