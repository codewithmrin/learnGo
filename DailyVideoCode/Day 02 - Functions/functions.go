package main

import "fmt"

// Function declaration
// func name(param1 type1) returnType (optional) {
// Function Body
//  }

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Print(sum(2, 3))
}