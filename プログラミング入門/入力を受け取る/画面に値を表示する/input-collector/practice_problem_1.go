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
		num, err := strconv.Atoi(scText);
		if err != nil {
			fmt.Printf("%v は整数ではありません。", scText)
			continue
		}
		if num < 0 {
			fmt.Printf("%v は正の整数ではありません。", scText)
			continue
		}

		L:
		scanner.Scan()
		scTexts := scanner.Text()
		scTextsArray := strings.Fields(scTexts)
		if num != len(scTextsArray) {
			fmt.Printf("最初に入力した%v個の数を入力してください。", num)
			goto L
		}

		var answer int
		for _, v := range scTextsArray {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("%v は正の整数ではありません。", v)
				goto L
			}
			answer += num
		}

		fmt.Println(answer)
		break
	}
}