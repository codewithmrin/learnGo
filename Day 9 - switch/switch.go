package main

import "fmt"

func main() {
	var userInpt string

	fmt.Println("Enter your Grade: ")

	fmt.Scan(&userInpt)

	switch userInpt {
	case "A":
		fmt.Println("You scored above 80%")
	case "B":
		fmt.Println("You scored above 60%")
	case "C":
		fmt.Println("You scored above 40%")
	case "F":
		fmt.Println("You scored below 40 and you have to repeat the subject")

	default:
		fmt.Println("Grade not Identified.")
	}
}
