package main

import (
    "fmt"
    "net"
    "net/rpc"
)

type MathService struct{}

func (m *MathService) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

type Args struct {
    A, B int
}

func main() {

    mathService := new(MathService)
    rpc.Register(mathService)


    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer listener.Close()

    fmt.Println("RPC server listening on port 1234...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }
        go rpc.ServeConn(conn)
    }
}
