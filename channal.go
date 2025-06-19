package main

import (
    "fmt"
    "time"
)

func sayHello(ch chan string) {
    time.Sleep(1 * time.Second) // simulate delay
    ch <- "Hello from Goroutine!" // send message into channel
}

func main() {
    ch := make(chan string) // create a channel of string type

    go sayHello(ch) // start goroutine

    message := <-ch // receive message from channel
    fmt.Println("Received:", message)
}
