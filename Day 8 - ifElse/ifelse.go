package main

import "fmt"

func main() {
  var num int
  fmt.Print("Enter a number: ")
  fmt.Scan(&num)
  
  if num > 10 {
    fmt.Println("Your number is greater than 10")
  } else if num > 5 {
    fmt.Println("Your number is between 5 and 10")
  } else {
    fmt.Println("Your number is less than or equal to 5")
  }
}
