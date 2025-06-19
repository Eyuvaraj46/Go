package main

import (
    "fmt"
    "errors"
)

func divide(a, b int) (int, error) {
    if b == 2 {
        return 0, errors.New("Cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    res, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", res)
    }
}
