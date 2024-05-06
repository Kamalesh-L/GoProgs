package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const maxCustomers = 10  // Total number of customers who will visit the barber shop
const waitingRoomSize = 3 // Number of seats in the waiting room

// Barber function
func barber(shop chan int, wg *sync.WaitGroup) {
    for {
        customer, ok := <-shop // Wait for a customer or close of channel
        if !ok {
            fmt.Println("All customers served. The barber goes home.")
            wg.Done()
            return
        }
        fmt.Printf("Barber is cutting hair of customer %d\n", customer)
        time.Sleep(time.Duration(rand.Intn(500)+100) * time.Millisecond) // Time to cut hair
        fmt.Printf("Barber finished with customer %d\n", customer)
    }
}

// Customer function
func customer(id int, shop chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Customer %d arrives at the barber shop.\n", id)

    select {
    case shop <- id: // Sit on a waiting room chair
        fmt.Printf("Customer %d is waiting in the waiting room.\n", id)
    default:
        fmt.Printf("Customer %d found no empty chair and leaves.\n", id)
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    var wg sync.WaitGroup
    shop := make(chan int, waitingRoomSize)

    wg.Add(1)
    go barber(shop, &wg)

    for i := 1; i <= maxCustomers; i++ {
        wg.Add(1)
        go customer(i, shop, &wg)
        time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond) // Interval between customers arriving
    }

    wg.Wait()
    close(shop) // No more customers will come, close the shop
}
