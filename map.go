package main
import "fmt"
func main(){
	fmt.Println("learn abt map in golang")
	var languages := make(map[string]string)
	var languages["JS"]="javascript"
	var languages["RB"]="ruby"
	var languages["PY"]="python"
	fmt.Println("list of all languages:",languages)
	fmt.Println("js shorts for:",languages["JS"])


	}	

