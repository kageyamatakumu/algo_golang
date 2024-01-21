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
		if num < 0 {
			fmt.Printf("%vは正の整数ではありません。\n", num)
			continue
		}

		L:
		scanner.Scan()
		scTexts := scanner.Text()
		scTextsArray := strings.Fields(scTexts)
		if num != len(scTextsArray) {
			fmt.Printf("最初に入力した%v個の数を入力してください。\n", num)
			goto L
		}

		for _, v := range scTextsArray {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("%v は正の整数ではありません。\n", v)
				goto L
			}
			fmt.Printf("%d\n", num % 10)
		}
		break
	}
}