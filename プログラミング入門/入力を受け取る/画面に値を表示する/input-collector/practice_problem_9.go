package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		scText := scanner.Text()
		num, err := strconv.Atoi(scText)
		if err != nil {
			fmt.Printf("%vは正の整数ではありません。\n", scText)
			continue
		}
		if num < 0 {
			fmt.Printf("%vは正の整数ではありません。\n", scText)
			continue
		}

		var contain string
		for i := 0; i < num; i++ {
			L:
			scanner.Scan()
			scText := scanner.Text()
			if scText == " " {
				fmt.Println("から文字が入力されました")
				goto L
			}
			contain += scText[:1]
		}

		fmt.Printf("%v\n", contain)
		break
	}
}