package main

import "fmt"

func main() {
    // Declaring and initializing an array
    arr := [5]uint8{1, 2, 3, 4, 5}

    // Accessing elements
    fmt.Println("Element at index 0:", arr[0]) // Outputs: 1

    // Looping through the array
    for i, v := range arr {
        fmt.Println("Index:", i, "Value:", v)
    }

    // Array length
    fmt.Println("Length of the array:", len(arr)) // Outputs: 5

    // Multi-dimensional array
    matrix := [2][2]int{{1, 2}, {3, 4}}
    fmt.Println("2D Array:", matrix)
}
