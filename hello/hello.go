package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//message, err := greetings.Hello("Rafal")
	// If an error, print it and exit program
	//if err != nil {
	//	log.Fatal(err)
	//}

	// If no error print message to the console
	//fmt.Println(message)

	names := []string{"Rafal", "Sabine", "Daria"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

}
