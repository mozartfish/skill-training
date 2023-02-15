package main

// The following code is from Rob Pike's 2012 Google I/0 talk
// explaining how concurrency is not parallelism
import (
	"fmt"
	"math/rand"
	"time"
)

// Restoring Sequencing
type Message struct {
	str  string
	wait chan bool
}

// Boring function runs on forever, like a boring party guest
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// Boring function with channels
// Channels connect the main and boring goroutines so they can communicate
func boringChannels(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// Rob Pike's Go concurrency patterns
// Generator: function that returns a channel
// Channels are first-class values, just like strings or integers
func boringGenerator(msg string) <-chan string { // Returns receive-only channels of strings
	c := make(chan string)
	go func() { // launch the goroutine from inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c // Return the channel to the caller
}

// Wait Channels and Restoring Sequencing Pattern
func boringGeneratorWait(msg string) <-chan Message { // Returns receive-only channel of Messages
	c := make(chan Message)
	waitForIt := make(chan bool) // Shared between all messages
	go func() {                  // launch the goroutine from inside the function
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c

}

// Multiplexing
// These programs make Joe and Ann count in lockstep
// We can instead use a fan-in function to let whosoever is ready to talk
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func fanInWait(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c

}

// Rob Pike Concurrency Definition
// concurrency is the composition of independently executing computations
// Concurrency is a way to structure software particularly as a way to write
// clean code that interacts well with the world
// concurrency is a model for software construction
// channels are first class citizens
// erlang -> writing to a file, go channels -> writing to a file descriptor

// Channels
// Declaring and intializing
// var c chan int // channels are typed
// c = make(chan int)
// c := make(chan int)
// Sending on a channel
// c <- 1
// Receiving from a channel
// The 'arrow' indicates the direction of data flow
// value = <-c

func main() {
	// The go statement runs the function as usual but doesn't make the caller wait.
	// It launches a goroutine - functionality is analogous to the & on the end of a shell command

	// Generator pattern
	// boring function returns a channel that lets us communicate with the boring
	// service it provides.
	// c := boringGenerator("boring!") // Function returning a channel
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("You say: %q\n", <-c)
	// }
	// fmt.Println("You're boring; I'm leaving.")

	// Channels as a handle on a service
	// Our boring function returns a channel that lets us communicate
	// with the boring service it provides.
	// We can have more instances of the service
	// joe := boringGenerator("Joe")
	// ann := boringGenerator("Ann")
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(<-joe)
	// 	fmt.Println(<-ann)
	// }
	// fmt.Println("You're both boring; I'm leaving.")

	// Multiplexing Pattern
	// These programs make Joe and Ann count in lockstep
	// We can instead use a fan-in function to let whosoever is ready talk
	// c := fanIn(boringGenerator("Joe"), boringGenerator("Ann"))
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-c)
	// }
	// fmt.Println("You're both boring; I'm leaving")

	// Restoring Sequencing
	// Send a channel on a channel making goroutine wait its turn
	// Receive all messages, then enable them again by sending on a private channel
	// First we define a message type that contains a channel for the reply
	// Each speaker must wait for a go-ahead
	c := fanInWait(boringGeneratorWait("Joe"), boringGeneratorWait("Ann"))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	// Wait Channels
	// Select Statement
	// FanIn Statement
	// Timeout
	// Communication
	// Synchronization
	// Chinese Whispers
	// Fake Search
	// Timeout Pattern
	// Replication Pattern
	// Summary
	// Other Examples
	// Conclusion
	// Boring goroutine channels
	// c := make(chan string)
	// go boringChannels("boring!", c)
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value
	// }
	// fmt.Println("You're boring; I'm leaving.")

	// Boring goroutine
	// go boring("boring!")
	// fmt.Println("I'm listening")
	// time.Sleep(2 * time.Second)
	// fmt.Println("You're boring; I'm leaving.")

	// Boring function
	// boring("boring!")

	// anonymous functions in go
	// func() {
	// 	fmt.Println("Hello World! This is an anonymous function")
	// }()
	// fmt.Println("Hello, World!") // for testing
}
