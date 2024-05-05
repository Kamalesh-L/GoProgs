package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    go startServer()

    // connect to the server
    conn, _ := net.Dial("tcp", "localhost:8081")

    for {
        // read in input from stdin
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Text to send: ")
        text, _ := reader.ReadString('\n')

        // send to the server
        fmt.Fprintf(conn, text+"\n")
    }
}

func startServer() {
    fmt.Println("Starting server...")

    // listen on a port
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer ln.Close()

    for {
        // accept a connection
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err.Error())
            return
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    for {
        // read from the connection
        message, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println("Client disconnected")
            return
        }

        // print the message
        fmt.Print("Message Received:", string(message))
    }
}