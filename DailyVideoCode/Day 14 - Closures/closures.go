package main

import "fmt"

func even_numbers() func() int {
	i := 0
	return func() int {
		i += 2
		return i
	}
}

func main() {
	next_even := even_numbers()
	fmt.Println(next_even())
	fmt.Println(next_even())
	fmt.Println(next_even())
	fmt.Println(next_even())
}
