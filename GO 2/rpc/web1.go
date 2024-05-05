package main

import (
    "net"
    "net/rpc"
    "log"
)

type Server struct{}

func (s *Server) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

type Args struct {
    A, B int
}

func main() {
    rpc.Register(new(Server))
    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("Listener error: ", err)
    }
    rpc.Accept(listener)
}