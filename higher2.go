package main
import "fmt"

func call() func(x int, y int){
	return add
}

func add(x int, y int){
	z := x + y
	fmt.Println(z)
}

func main(){
	sum := call() // Call the call function, which returns the add function
	sum(5, 10) // Call the returned add function with arguments 5 and 10
}