package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel to store Fibonacci numbers
	ch := make(chan int, 10) // Use a buffer size of 10 for demonstration

	// Goroutine to generate Fibonacci numbers and write them to the channel
	go func() {
		defer close(ch) // Close the channel when done producing Fibonacci numbers
		fibonacci(ch)
	}()

	// Goroutine to read and consume Fibonacci numbers from the channel
	for num := range ch {
		fmt.Println(num)
	}
}

// Function to generate Fibonacci numbers and write them to a channel
func fibonacci(ch chan<- int) {
	a, b := 0, 1
	for i := 0; i < 10; i++ { // Generate first 10 Fibonacci numbers
		ch <- a // Write Fibonacci number to the channel
		a, b = b, a+b // Update Fibonacci sequence
	}
}
