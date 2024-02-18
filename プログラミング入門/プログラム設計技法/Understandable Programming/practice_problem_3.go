package main

import "fmt"

func main() {
	var radius int = 5
	var Pi float32 = 3.14

	fmt.Println(float32(2) * float32(radius) * Pi)
	fmt.Println(float32(radius) * float32(radius) * Pi)
	fmt.Println(float32(radius) * float32(radius) * Pi * float32(10))
	fmt.Println(float32(4) * float32(radius) * float32(radius) * Pi)
	fmt.Println(float32(4) / 3.0 * float32(radius) * float32(radius) * float32(radius) * Pi)

}



