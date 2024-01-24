package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"Gladys",
		"John",
		"Peace",
	}

	ages := []string{
		"18",
		"24",
		"56",
	}

	// Get a greeting message and print it here.
	messages, err := greetings.Hellos(names, ages)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
