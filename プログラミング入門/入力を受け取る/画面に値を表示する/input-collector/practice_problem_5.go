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
			fmt.Printf("%vは正の整数ではありません。\n", num)
			continue
		}

		L:
		scanner.Scan()
		scTexts := scanner.Text()
		scTextsArray := strings.Fields(scTexts)
		if num != len(scTextsArray) {
			fmt.Printf("最初に入力した%d個の数を入力åしてください。\n", num)
			goto L
		}

		for i := len(scTextsArray)-1; i >= 0; i-- {
			num, err := strconv.Atoi(scTextsArray[i])
			if err != nil {
				fmt.Printf("%v は正の整数ではありません。\n", num)
				goto L
			}
			fmt.Printf("%d\n", num)
		}
		break
	}
}

