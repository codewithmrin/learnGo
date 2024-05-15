package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println("add: 3, 4 =", add(3, 4))
	fmt.Println("sum: 3, 4 =", sum(3, 4))
	fmt.Println("sum: 3, 4, 5 =", sum(3, 4, 5))
	arr := []int{2, 2, 3, 4}
	fmt.Println("sum:", arr, "=", sum(arr...))
}
