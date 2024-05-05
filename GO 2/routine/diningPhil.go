package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numPhilosophers = 5
	eatLimit        = 3 // Number of times each philosopher can eat
)

var (
	chopsticks [numPhilosophers]sync.Mutex
	eatCount   [numPhilosophers]int // Track eat count for each philosopher
	wg         sync.WaitGroup
)

func main() {
	wg.Add(numPhilosophers)

	for i := 0; i < numPhilosophers; i++ {
		go philosopher(i)
	}

	wg.Wait()
	fmt.Println("All philosophers are satisfied. Program ends.")
}

func philosopher(id int) {
	left, right := id, (id+1)%numPhilosophers

	for eatCount[id] < eatLimit {
		fmt.Printf("Philosopher %d is thinking\n", id)
		time.Sleep(time.Second)

		fmt.Printf("Philosopher %d is hungry\n", id)

		chopsticks[left].Lock()
		fmt.Printf("Philosopher %d has acquired left chopstick %d\n", id, left)

		chopsticks[right].Lock()
		fmt.Printf("Philosopher %d has acquired right chopstick %d\n", id, right)

		fmt.Printf("Philosopher %d is eating\n", id)
		time.Sleep(time.Second)

		fmt.Printf("Philosopher %d has finished eating\n", id)

		chopsticks[right].Unlock()
		fmt.Printf("Philosopher %d has released right chopstick %d\n", id, right)

		chopsticks[left].Unlock()
		fmt.Printf("Philosopher %d has released left chopstick %d\n", id, left)

		eatCount[id]++
	}

	fmt.Printf("Philosopher %d is satisfied\n", id)
	wg.Done()
}
