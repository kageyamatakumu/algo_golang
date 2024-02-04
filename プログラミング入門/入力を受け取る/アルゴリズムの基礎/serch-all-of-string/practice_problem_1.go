package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scText := scanner.Text()

	scanner.Scan()
	checkChar := scanner.Text()

	flug := strings.ContainsAny(scText, checkChar)

	if flug {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
