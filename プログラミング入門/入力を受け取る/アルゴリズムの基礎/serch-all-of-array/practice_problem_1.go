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
		scManyTexts := scanner.Text()
		scManyTextsArray := strings.Fields(scManyTexts)
		if len(scManyTextsArray) != scTwoNumber[0] {
			fmt.Printf("%dと入力された個数が一致しません。\n", scTwoNumber[0])
			goto L2
		}

		var answer string
		for _, v := range scManyTextsArray {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("%vは正の整数ではありません。\n", v)
				goto L2
			}
      if num == scTwoNumber[1] {
        answer = "Yes"
        break
      } else {
        answer = "No"
      }
		}
		fmt.Printf("%v\n", answer)
    break
	}
}