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

func (m *MathService) Divide(args *Args, reply *float64) error {
    if args.B == 0 {
        return fmt.Errorf("division by zero")
    }
    *reply = float64(args.A) / float64(args.B)
    return nil
}

type GreetingService struct{}

func (g *GreetingService) Greet(name string, reply *string) error {
    *reply = "Hello, " + name + "!"
    return nil
}

type Args struct {
    A, B int
}

func main() {
    mathService := new(MathService)
    greetingService := new(GreetingService)

    rpc.Register(mathService)
    rpc.Register(greetingService)

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
