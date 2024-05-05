package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 0; i < 10; i++ {
        time.Sleep(1 * time.Second)
        fmt.Println(i)
    }
}

func printLetters() {
    for i := 'a'; i < 'a'+10; i++ {
        time.Sleep(1 * time.Second)
        fmt.Printf("%c ", i)
    }
}

func main() {
    go printNumbers()
    go printLetters()
    time.Sleep(11 * time.Second)
}