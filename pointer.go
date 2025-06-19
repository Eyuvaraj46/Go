package main
import "fmt"

func main(){
	x:=10
	p:=&x
	fmt.Println("value of x:",x)
	fmt.Println("addres of x",p)
	fmt.Println("value of p",*p)

	*p=20
	fmt.Println("update of x:",x)


}
