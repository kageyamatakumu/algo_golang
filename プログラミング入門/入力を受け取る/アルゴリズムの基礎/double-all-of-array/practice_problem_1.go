package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isPrime関数は与えられた数が素数かどうかを判定します。
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= n / 2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

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

		scanner.Scan()
		scTexts := scanner.Text()
		scTextsArray := strings.Fields(scTexts)
		if len(scTextsArray) != num {
			fmt.Println("最初に入力した数字の数と一致しません。もう一度やり直してください。")
			continue
		}

		primeNumber := 0

		for _, v := range scTextsArray {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("%vは正の整数ではありません。\n", v)
				continue
			}
			if isPrime(num) {
				primeNumber++
			}
		}

		fmt.Printf("%d\n", primeNumber)
		break
	}
}
