package main

import "fmt"

func main() {
	var grade string

	fmt.Print("Enter your grade: ")
	fmt.Scan(&grade)
	switch grade {
	case "A":
		fmt.Println("You scored above 80%")
	case "B":
		fmt.Println("You scored above 60%")
	case "C":
		fmt.Println("You scored above 40%")
	case "F":
		fmt.Println("You did not pass this course")
	default:
		fmt.Println("Grade Unnknown")
	}
}
