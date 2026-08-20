package main
import "fmt"

// The init function is a special function in Go that is automatically executed before the main function. It is used to initialize variables or perform setup tasks before the program starts executing. The init function can be defined in any package, and it can be used to set up package-level variables or perform other initialization tasks.

var a = 10

func main() {
	fmt.Println("The a is: ", a)
}

func init(){
	fmt.Println("The a is: ", a)
	a = 20
}