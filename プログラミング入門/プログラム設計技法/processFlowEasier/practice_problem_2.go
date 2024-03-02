package main

import "fmt"

func main() {
	fmt.Println(opposite_side(1))
	fmt.Println(opposite_side(7))
	fmt.Println(opposite_side(3))
	fmt.Println(opposite_side(2))
	fmt.Println(opposite_side(0))
	fmt.Println(opposite_side(5))
}

func opposite_side(num int) int {
	if num < 1 || num > 6 {
		return -1
	}

	return 7 - num
}