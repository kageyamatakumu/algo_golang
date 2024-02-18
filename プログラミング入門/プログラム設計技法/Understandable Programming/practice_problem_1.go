package main

import "fmt"

func main(){
	var answer int = 1;

	for i := 1; i <= 10; i++ {
		answer *= i
	}

	fmt.Println(answer)
}