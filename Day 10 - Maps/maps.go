package main

import "fmt"

func main() {
	var map1 = make(map[string]int)

	map1["key1"] = 1
	map1["key2"] = 2

	fmt.Println(map1["key1"])
	fmt.Println(map1["key2"])
	fmt.Println(map1["Nokey"])
}
