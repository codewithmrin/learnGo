package main

import "fmt"

func main() {
	var a = 10
	for a > 0 {
		fmt.Println("a", a)
		a -= 2
	}

	var b int
	for b = 0; b < 10; b += 1 {
		fmt.Println("b", b)
	}
}
