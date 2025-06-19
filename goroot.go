package main

import (
    "fmt"
    "time"
)

// Print numbers from 1 to 5
func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println("Number:", i)
        time.Sleep(300 * time.Millisecond)
    }
}

// Print letters A to E
func printLetters() {
    for _, letter := range []string{"A", "B", "C", "D", "E"} {
        fmt.Println("Letter:", letter)
        time.Sleep(300 * time.Millisecond)
    }
}

func main() {
    go printNumbers() // run in background
    go printLetters() // run in background

    // Wait for both goroutines to finish
    time.Sleep(2 * time.Second)

    fmt.Println("Main function done.")
}
