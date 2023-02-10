package main

import "fmt"
import "example.com/greetings"
import "log"

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	// message, err := greetings.Hello("")
	// message, err := greetings.Hello("James Bond")

	// A slice of names.
	names := []string{"James Bond", "Bruce Wayne", "Daft Punk"}

	// Request greeting message for the names
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console
	// and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console
	fmt.Println(messages)
}
