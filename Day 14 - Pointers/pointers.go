package main

import "fmt"

func main() {
	var a int

	a = 10
	ptr := &a
	fmt.Println(ptr)
	fmt.Println(*ptr)

	// Day 15 - Using Pointers
	fmt.Println("Before change val:", a)

	change_val(a)
	fmt.Println("after change val:", a)

	change_ref(&a)
	fmt.Println("after change ref:", a)

}

func change_val(val int) {
	fmt.Println("before change val:", val)
	val = 15
}

func change_ref(val *int) {
	fmt.Println("before change ref:", *val)
	*val = 15
}
