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

	scanner.Scan()

	for {
		scText := scanner.Text()
		num, err := strconv.Atoi(scText)
		if err != nil {
			fmt.Printf("%vは正の整数ではありません。\n", scText)
			continue
		}


		L:
		scanner.Scan()
		scTexts := scanner.Text()
		scTextsArray := strings.Split(scTexts, " ")

		if num != len(scTextsArray) {
			fmt.Printf("%dと入力された個数が一致しません。\n", len(scTextsArray))
			goto L
		}

		var min int
		for i, v := range scTextsArray {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("%vは正の整数ではありません。\n", v)
				goto L
			}

			if i == 0 {
				min = num
			}

			if min > num {
				min = num
			}
		}

		fmt.Printf("%d\n", min)
		break
	}

}