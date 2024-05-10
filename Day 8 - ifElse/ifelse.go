package main

import "fmt"

func main() {
	var marks int
	fmt.Print("Enter your Marks: ")
	fmt.Scan(&marks)
	if marks > 80 {
		fmt.Println("Grade A")
	} else if marks > 60 {
		fmt.Println("Grade B")
	} else if marks > 40 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Fail...")
	}
}