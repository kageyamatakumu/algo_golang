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

		scNum, err := strconv.Atoi(scText)
		if err != nil {
			fmt.Printf("%vは正の整数ではありません。\n", scNum)
			continue
		}

		scanner.Scan()

		L:
		scTexts := scanner.Text()
		scTextsArray := strings.Fields(scTexts)

		if len(scTextsArray) != scNum {
			fmt.Printf("%dと入力された個数が一致しません。\n", scNum)
			goto L
		}

		var max int
		for _, v := range scTextsArray {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("%vは正の整数ではありません。\n", v)
				goto L
			}

			if max < num {
				max = num
			}
		}

		fmt.Printf("%d\n", max)
		break
	}
}