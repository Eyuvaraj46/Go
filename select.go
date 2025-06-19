package main

import (
    "fmt"
    "time"
)

func fastChannel(ch chan string) {				
    time.Sleep(1 * time.Second)
    ch <- "Fast channel finished!"
}

func slowChannel(ch chan string) {
    time.Sleep(2 * time.Second)
    ch <- "Slow channel finished!"
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go fastChannel(ch1)
    go slowChannel(ch2)

    // use select to wait for whichever sends first
	//It lets you wait on multiple channel operations and runs whichever one is ready first.
	//select in Go is like a switch statement — but for channels.
    select {
    case msg1 := <-ch1:
        fmt.Println("Received:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Received:", msg2)
    default:
        fmt.Println("No channel is ready.")
    }
}
