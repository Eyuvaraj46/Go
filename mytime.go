package main

import ("fmt"
		"time"
) 


func main() {
	fmt.Println("welcome to time study of go ")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //this is a format of go
	
	createdDate :=time.Date(2020,time.December,12,23,23,0,0,time.UTC) //this is a method to create your own time etc..
	
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))

	


}
