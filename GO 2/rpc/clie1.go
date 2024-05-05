package main

import (
    "fmt"
    "net/rpc"
)

type Args struct {
    A, B int
}

func main() {
    client, err := rpc.Dial("tcp", "localhost:1234")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer client.Close()

    var multiplyResult int
    var divideResult float64
    var greetResult string

    args := &Args{4, 5}
    err = client.Call("MathService.Multiply", args, &multiplyResult)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Multiplication Result:", multiplyResult)

    args = &Args{10, 2}
    err = client.Call("MathService.Divide", args, &divideResult)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Division Result:", divideResult)

    err = client.Call("GreetingService.Greet", "Kamalesh", &greetResult)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Greeting:", greetResult)
}
