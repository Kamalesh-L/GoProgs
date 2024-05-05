package main

import (
    "net/rpc"
    "log"
    "fmt"
)

type Args struct {
    A, B int
}

func main() {
    client, err := rpc.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing error: ", err)
    }
    args := &Args{7,8}
    var reply int
    err = client.Call("Server.Multiply", args, &reply)
    if err != nil {
        log.Fatal("rpc error: ", err)
    }
    fmt.Printf("Result: %d*%d=%d\n", args.A, args.B, reply)
}