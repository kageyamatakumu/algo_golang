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

	scanner.Scan()
	scText2 := scanner.Text()

	if scText == scText2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}