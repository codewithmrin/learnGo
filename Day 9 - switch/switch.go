package main

import "fmt"

func main() {
  var myNum int
  fmt.Print("Enter a number: ")

  fmt.Scan(&myNum)

  switch myNum {
    case 10: fmt.Println("Your number is 10")
    case 20: fmt.Println("Your number is 20")
    case 30: fmt.Println("Your number is 30")
    default: fmt.Printf("Your number is not 10, 20, or 30. But it is %d\n", myNum)
  }
}
