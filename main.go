package main

import (
	"fmt"
	"log"

	greetings "github.com/atilasantos/go-greetings"
)

func main() {

	log.SetFlags(0)
	log.SetPrefix("Greetings: ")

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Message:" + message)

}
