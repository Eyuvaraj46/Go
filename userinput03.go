package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	welcome :="welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our pizza:")
	//comm ok //err ok
	input, _ := reader.ReadString('\n')//getting input form the user
	fmt.Println("thanks for rating:", input)
	fmt.Printf("Type of this rating is %T", input)


}