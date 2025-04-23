package main
import "fmt"

func main(){
	fmt.Println("welcome to slice opration")

	var fruitlist = []string{"apple","peach","jack fruit"}
	fmt.Printf("type of fruitlist is%T\n",fruitlist)

	fruitlist = append(fruitlist , "mango","banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:])
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)

	highscores := make([]int,4)

	highscores[0] = 111
	highscores[1] = 968
	highscores[2] = 546
	highscores[3] = 855
	//highscores[4] = 234 its will not works 

	highscores = append(highscores,555,222,666 )

	fmt.Println(highscores)



}