package main
import (
    "fmt"
    "time"
)

func greet() {
    for i := 1; i <= 3; i++ {
        fmt.Println("Hello from Goroutine")
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    go greet()
    fmt.Println("Main is running")
    time.Sleep(2 * time.Second) //go asking greet to wait 2 second
}
