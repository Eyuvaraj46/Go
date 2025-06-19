package main

import ("fmt"
		"errors")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("connot divide by zero")

	}
	return a / b, nil //no err

}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result)
	}

	result, err = divide(5, 0)
	if err != nil {
		fmt.Println("error", err)

	} else {
		fmt.Println("result:", result)
	}
}
