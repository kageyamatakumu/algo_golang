package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var text string
	for i := 0; i < 3; i++ {
		text += scanner.Text()
	}

	fmt.Println(text)
}