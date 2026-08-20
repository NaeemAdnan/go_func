package main
import "fmt"

// The init function is a special function in Go that is automatically executed before the main function. It is used to initialize variables or perform setup tasks before the program starts executing. The init function can be defined in any package, and it can be used to set up package-level variables or perform other initialization tasks.
// Standard function is a regular function that is defined with a name and can be called from other parts of the program. It can take parameters, return values, and perform various operations. Standard functions are the building blocks of Go programs and are used to encapsulate logic and functionality.

var (
	a = 10
	b = 20
)

func add() int {
	z := a + b
	return z
}

func main() {
	fmt.Println(add())
}