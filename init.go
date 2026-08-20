package main

import "fmt"

var a = 10

func main() {
	fmt.Println("The a is: ", a)
}

func init(){
	fmt.Println("The a is: ", a)
	a = 20
}