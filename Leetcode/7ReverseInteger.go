package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Print("Enter number to reverse:")
	fmt.Scan(&num)
	fmt.Println(reverse(num))
}

func reverse(x int) int {
	op := 0
	in := x
	s := 1
	if in < 0 {
		in = -in
		s = -1
	}
	for in > 0 {
		digit := in % 10
		in = in / 10
		op = op*10 + digit
		if op > math.MaxInt32 || op < math.MinInt32 {
			return 0
		}
	}
	return op * s
}
