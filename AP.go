package main
import "fmt"

func add(a int, b int){ // Define a function named add that takes two integer parameters a and b
	c := a + b
	fmt.Println(c)
}

func main(){
	add(6,7) // Call the add function. Arguments 6 and 7 are passed to the function, which adds them and prints the result (13).
}

// Parametes vs Arguments: Parameters are the variables defined in the function signature (e.g., a and b in the add function). Arguments are the actual values passed to the function when it is called (e.g., 6 and 7 in the main function).

// Arguments vs Parameters: Arguments are the actual values passed to the function when it is called (e.g., 6 and 7 in the main function). Parameters are the variables defined in the function signature (e.g., a and b in the add function).