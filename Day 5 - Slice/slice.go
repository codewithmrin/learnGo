package main

import "fmt"

// Slicing and Array Concatenation
func main() {
	arr1 := []int{2, 3, 4, 5}
	// arrayName[start:end]
	// we get slice from index start to end - 1
	mySlice := arr1[1:3]

	arr2 := append(arr1, mySlice[0])

	arr3 := append(arr1, 34, 35, 34, 23, 12, 232, 43, 23)
	
	fmt.Println(arr1)
	fmt.Println(mySlice)
	fmt.Println(arr2)
	fmt.Println(arr3)
}