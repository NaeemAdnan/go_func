package main

import "fmt"

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