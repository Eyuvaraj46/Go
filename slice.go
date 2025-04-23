package main
import "fmt"

func main() {
    arr := []int8{1, 2, 3, 4, 5}
    slice := arr[1:4]
	fmt.Println(slice)
    slice[0] = 96 + 3 // Modifies arr[1]
    
    fmt.Println(arr)   // Output: [1 99 3 4 5]
    fmt.Println(slice) // Output: [99 3 4]

}
