package main

import "fmt"

func main() {
	fmt.Println(checkMental(3))
	fmt.Println(checkMental(-1))
	fmt.Println(checkMental(0))
	fmt.Println(checkMental(1))
	fmt.Println(checkMental(-13))
}

func checkMental(num int) string {
	if num > 0 {
		return "positive"
	} else if num < 0 {
		return "negative"
	} else {
		return "zero"
	}
}