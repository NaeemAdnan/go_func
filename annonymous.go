package main
import "fmt"

// Annonymous function is a function without a name. It can be defined and called at the same time. In Go, anonymous functions are often used as closures, which means they can capture and use variables from their surrounding scope.
// IIFE (Immediately Invoked Function Expression) is a common pattern in JavaScript, but it can also be used in Go. An IIFE is a function that is defined and immediately called, allowing you to create a new scope and avoid polluting the global namespace.

var(
	a = 10
	b = 20
)

func main(){
	func(){
		c := a+b
		fmt.Println("The c is: ", c)
	} ()



}