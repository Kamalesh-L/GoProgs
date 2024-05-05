package main

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
	"sync"
)

// Function to sum numbers in parallel
func parallelSumFromInput() int {
	var wg sync.WaitGroup
	numbers := make(chan int)
	results := make(chan int)

	// Goroutine to read numbers from input
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(numbers)

		fmt.Println("Enter numbers (one per line), type 'done' when finished:")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			if input == "done" {
				break
			}

			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("Invalid input: %v\n", err)
				continue
			}

			numbers <- num
		}
	}()

	// Worker goroutine to compute partial sums
	const numWorkers = 4
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			partialSum := 0
			for num := range numbers {
				partialSum += num
			}
			results <- partialSum
		}()
	}

	// Goroutine to aggregate partial sums and compute final result
	go func() {
		wg.Wait()
		close(results)
	}()

	totalSum := 0
	for partialSum := range results {
		totalSum += partialSum
	}

	return totalSum
}

func main() {
	totalSum := parallelSumFromInput()
	fmt.Printf("Total sum: %d\n", totalSum)
}
