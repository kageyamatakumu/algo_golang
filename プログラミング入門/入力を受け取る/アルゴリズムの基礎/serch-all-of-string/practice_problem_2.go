package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scText := scanner.Text()

	length := len(scText)
	for i := 0; i < len(scText); i++ {
		if scText[i] == scText[len(scText) - 1 - i] {
			length--
		}
	}

	if length == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}