package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
    // connect to the server
    conn, _ := net.Dial("tcp", "localhost:8080")

    go listenForReply(conn)

    for {
        // read in input from stdin
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Text to send: ")
        text, _ := reader.ReadString('\n')

        // send to the server
        fmt.Fprintf(conn, text+"\n")
    }
}

func listenForReply(conn net.Conn) {
    for {
        // listen for reply
        message, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Print("Message from server: "+message)
    }
}