package main

import "fmt"

func main() {
	fmt.Println(quadrant(1, 2))
	fmt.Println(quadrant(-1,-3))
	fmt.Println(quadrant(-2,1))
	fmt.Println(quadrant(-3, -2))
	fmt.Println(quadrant(2, -1))
}

func quadrant(x int, y int) int {
	if x > 0 && y > 0 {
		return 1
	} else if x < 0 && y > 0 {
		return 2
	} else if x < 0 && y < 0 {
		return 3
	} else if x > 0 && y < 0 {
		return 4
	} else {
		return 0
	}
}
