package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(20, 3)
	fmt.Println("sum:", result)

}
