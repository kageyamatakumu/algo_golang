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
		scNumber, err := strconv.Atoi(scText)
		if err != nil {
			fmt.Printf("%vは数値ではありません。\n", scText)
			continue
		}

		L:
		scanner.Scan()
		scTexts := scanner.Text()
		scTextsArray := strings.Fields(scTexts)

		if scNumber != len(scTextsArray) {
			fmt.Printf("%dと入力された個数が一致しません。\n", scNumber)
			goto L
		}

		var positiveInteger int
		for _, v := range scTextsArray {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("%vは数値ではありません。\n", v)
				goto L
			}

			if num > 0 {
				positiveInteger++
			}
		}

		fmt.Printf("%d\n", positiveInteger)
		break
	}
}