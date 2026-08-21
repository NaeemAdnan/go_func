package main
import "fmt"

// function expression is a way to define a function without giving it a name. It allows you to create anonymous functions that can be assigned to variables or passed as arguments to other functions. Function expressions are useful for creating closures, callbacks, and higher-order functions.

func main(){

	add := func(a int, b int){
		c := a + b
		fmt.Println(c)
	}
	add(5, 3) // Call the anonymous function
}

func init(){
	fmt.Println("This is the init function. It runs before the main function.")
}