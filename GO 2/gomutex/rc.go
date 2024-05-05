// package main  
// import (  
//     "fmt"
//     "sync"
//     )
// var x  = 0  
// func increment(wg *sync.WaitGroup) {  
//     x = x + 1
//     wg.Done()
// }
// func main() {  
//     var w sync.WaitGroup
//     for i := 0; i < 500;i++ {
//         w.Add(1)        
//         go increment(&w)
//     }
//     w.Wait()
//     fmt.Println("final value of x", x)
// }
package main

import (
    "fmt"
    "sync"
)

var (
    x      = 0
    mutex  sync.Mutex // Mutex to synchronize access to the shared variable x
)

func increment(wg *sync.WaitGroup) {
    // Lock the mutex to ensure exclusive access to the shared variable x
    mutex.Lock()
    x = x + 1
    // Unlock the mutex once the modification is done
    mutex.Unlock()
    
    wg.Done()
}

func main() {
    var w sync.WaitGroup

    // Launch 500 goroutines, each incrementing the value of x
    for i := 0; i < 500; i++ {
        w.Add(1) // Add 1 to the WaitGroup for each goroutine
        go increment(&w) // Start a new goroutine to execute the increment function
    }

    w.Wait() // Wait for all goroutines to finish

    fmt.Println("final value of x", x) // Print the final value of x
}
