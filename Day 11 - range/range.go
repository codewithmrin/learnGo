package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	// Order is preserved in an array
	for _, number := range numbers {
		fmt.Printf("Value is %d\n", number)
	}

	lettersMap := map[string]int{"a": 1, "b": 2, "c": 3}
	// Order is not preserved in a map
	for key, value := range lettersMap {
		fmt.Printf("Key: %s Value: %d\n", key, value)
	}
}
