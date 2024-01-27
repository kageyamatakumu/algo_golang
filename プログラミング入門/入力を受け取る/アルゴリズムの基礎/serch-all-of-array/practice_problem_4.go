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

		L:
		scTwoTexts := scanner.Text()
		scTwoTextsArray := strings.Fields(scTwoTexts)

		if len(scTwoTextsArray) != 2 {
			fmt.Println("2つの整数を入力してください。")
			goto L
		}

		var scTwoNumber []int
		for _, v := range scTwoTextsArray {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("%vは正の整数ではありません。\n", v)
				goto L
			}

			scTwoNumber = append(scTwoNumber, num)
		}

		L2:
		scanner.Scan()
		scTexts := scanner.Text()
		scTextsArray := strings.Fields(scTexts)

		if len(scTextsArray) != scTwoNumber[0] {
			fmt.Printf("%dと入力された個数が一致しません。\n", scTwoNumber[0])
			goto L2
		}

		var answer int
		answer = 0
		for i, v := range scTextsArray {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("%vは正の整数ではありません。\n", v)
				goto L2
			}

			if num == scTwoNumber[1] {
				answer = i + 1
			}
		}

		if answer != 0 {
			fmt.Printf("%d\n", answer - 1)
			break
		} else {
			fmt.Println(-1)
			break
		}

	}
}