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

    args := &Args{4, 5}

    var result int
    err = client.Call("MathService.Multiply", args, &result)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Result:", result)
}
