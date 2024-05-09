package main

import "fmt"

func main() {
	var myNumber int

	fmt.Print("Enter a number: ")

	fmt.Scan(&myNumber)

	fmt.Printf("Your number is %d\n", myNumber)
}