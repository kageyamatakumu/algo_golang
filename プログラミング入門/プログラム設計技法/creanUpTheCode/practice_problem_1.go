package main

import (
	"fmt"
	"strconv"
)

func main() {
  for i := 1; i <= 15; i++ {
		fmt.Println(fizz_buzz(i))
	}
}

func fizz_buzz(num int) string {
  if num % 3 == 0 && num % 5 == 0 {
    return "FizzBuzz"
  } else if num % 3 == 0 {
    return "Fizz"
  } else if num % 5 == 0 {
    return "Buzz"
  } else {
    return strconv.Itoa(num)
  }
}