package main
import "fmt"

// Variable shadowing occurs when a variable declared within a certain scope (e.g., a function or block) has the same name as a variable declared in an outer scope. In such cases, the inner variable "shadows" the outer variable, making the outer variable inaccessible within that scope. This can lead to confusion and bugs if not handled carefully.

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