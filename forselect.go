package main

import (
	"fmt"
	"time"
)

func sendEven(ch chan int) {
	for i := 0; i <= 10; i += 2 {
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func sendOdd(ch chan int) {
	for i := 1; i <= 9; i += 2 {
		ch <- i
		time.Sleep(700 * time.Millisecond)
	}
	close(ch)
}

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)

	go sendEven(evenCh)
	go sendOdd(oddCh)

	// Keep track of closed channels
	evenClosed := false
	oddClosed := false

	for {
		if evenClosed && oddClosed {
			break // exit when both channels are closed
		}

		select {
		case val, ok := <-evenCh:
			if ok {
				fmt.Println("Even:", val)
			} else {
				evenClosed = true
			}

		case val, ok := <-oddCh:
			if ok {
				fmt.Println("Odd :", val)
			} else {
				oddClosed = true
			}
		}
	}

	fmt.Println("Finished receiving all values.")
}
