package main

import (
	"fmt"
	"log"

	greetings "github.com/atilasantos/go-greetings"
)

func main() {

	log.SetFlags(0)
	log.SetPrefix("Greetings: ")

	names := []string{"Juanito", "Pedrito", "Carlito"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	for name, message := range messages {
		fmt.Printf("The %v has the following message: %v\n", name, message)
	}

}
