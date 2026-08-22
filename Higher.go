package main
import "fmt"


func ProcessData(a int, b int, op func(p int, q int)){
	op(a, b)
}

func add(x int, y int){
	z := x + y
	fmt.Println(z)
}

func main(){
	ProcessData(5, 10, add) // Call ProcessData with the add function as an argument
}

func init(){
	fmt.Println("Higher-order functions are functions that can take other functions as arguments or return functions as results. They allow for more flexible and reusable code by enabling the creation of functions that can operate on different behaviors or operations.")
}