package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

		L:
		scanner.Scan()
		scTexts := scanner.Text()
		scTextsArray := strings.Fields(scTexts)

		if num != len(scTextsArray) {
			fmt.Printf("%dと入力された個数が一致しません。\n", len(scTextsArray))
			goto L
		}

		for i := 1; i <= 9; i++ {
			count := 0
			for _, v := range scTextsArray {
				num, err := strconv.Atoi(v)
				if err != nil {
					fmt.Printf("%vは正の整数ではありません。\n", v)
					goto L
				}

				if num == i {
					count++
				}
			}
			fmt.Printf("%d\n", count)
		}
		break
	}
}