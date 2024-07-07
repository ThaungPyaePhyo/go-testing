package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main() {

	log.SetPrefix("greetings: ")
    log.SetFlags(0)


	names := []string{"Gladys", "Samantha", "Darrin"}

	message, err := greeting.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
    fmt.Println(message)
}