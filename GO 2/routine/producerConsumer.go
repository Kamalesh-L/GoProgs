package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

const (
	bufferSize = 5   // Size of the buffer
	numProducers = 2 // Number of producer goroutines
	numConsumers = 3 // Number of consumer goroutines
	numItemsPerProducer = 5 // Number of items each producer produces
)

func main() {
	// Create a buffered channel to represent the shared buffer
	buffer := make(chan int, bufferSize)

	// Use a wait group to wait for all producers and consumers to finish
	var wgProducers, wgConsumers sync.WaitGroup
	wgProducers.Add(numProducers)
	wgConsumers.Add(numConsumers)

	// Start producer goroutines
	for i := 1; i <= numProducers; i++ {
		go producer(i, numItemsPerProducer, buffer, &wgProducers)
	}

	// Start consumer goroutines
	for i := 1; i <= numConsumers; i++ {
		go consumer(i, buffer, &wgConsumers)
	}

	// Wait for all producers and consumers to finish
	wgProducers.Wait()
	close(buffer) // Close the buffer channel to signal consumers to stop
	wgConsumers.Wait()
	fmt.Println("All producers and consumers have finished. Program ends.")
}

func producer(id, numItems int, buffer chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < numItems; i++ {
		item := rand.Intn(100) // Generate a random item
		fmt.Printf("Producer %d producing item %d\n", id, item)
		buffer <- item         // Add item to the buffer (blocking if buffer is full)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate production time
	}

	fmt.Printf("Producer %d finished producing\n", id)
}

func consumer(id int, buffer <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for item := range buffer {
		fmt.Printf("Consumer %d consuming item %d\n", id, item)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) // Simulate consumption time
	}

	fmt.Printf("Consumer %d finished consuming\n", id)
}
