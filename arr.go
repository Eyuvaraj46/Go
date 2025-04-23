package main

import "fmt"

func main(){
	fmt.Println("welcome to learn golangs")
	var fruitlist[4]string

	fruitlist[0]="apple"
	fruitlist[1]="orange"
	fruitlist[3]="banana"
	
	fmt.Println("fruit list  is",fruitlist)
	fmt.Println("fruit list  is",len(fruitlist))

	var veglist=[5]string{"potato","mushroom","carrot"}
	fmt.Println("veglist is",veglist)
	fmt.Println("veglist is ",len(veglist))





}