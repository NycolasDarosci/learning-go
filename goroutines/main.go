package main

import (
	"fmt"
	"time"
)

// what is the lifecycle of the app?
// the app ends when the main goroutine ends
func main() { // main goroutine

	var channel chan string
	go print("testing routine!", channel)
	//response := <-channel
	//channel := make(chan string)

	//go print("testing routine two!")

	// this print() is in the main go routine
	//print("testing routine three!")

	/*
		if put 'go' on the third call function
		the app will print nothing
		because the main go routine is killed
		and it must keep busy
		we can add for {} // infinite loop
		or time.Sleep()
		but the ideal is that don't kill the main routine process
	*/
	//go print("testing routine three!")
	//for {}
	//time.Sleep(time.Minute)
}

func print(message string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(800 * time.Millisecond)
	}
	channel <- "Done"
}
