package main
import "fmt"

var a = 10

func main(){
	fmt.Println("The a is: ", a)
	age := 30
	if age > 18{
		a = 47
		fmt.Println("The a is: ", a)
	}
	fmt.Println("The a is: ", a)
}