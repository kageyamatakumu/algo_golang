package main

import "fmt"

func main() {
	areaPrint(3)
	areaPrint(5)
	areaPrint(10)
	areaPrint(15)
}

func areaPrint(radius int) {
	var Pi float32 = 3.14
	fmt.Println(float32(radius) * float32(radius) * Pi)
}