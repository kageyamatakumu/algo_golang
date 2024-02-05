package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		scText := scanner.Text()

		if scText == "" {
			fmt.Println("空文字列は回文ではありません。")
			continue
		}

		count := 0
		for i := 0; i < len(scText) - 1; i++ {
			if scText[i] == scText[i+1] {
				count++
			}
		}

		fmt.Printf("%d\n", count)
		break
	}
}