package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		N := scanner.Text()
		if !isNotEmpty(N) || !sameLength(num, N) {
			fmt.Println("指定した文字数の文字を入力してください。")
			goto L
		}

		L2:
		scanner.Scan()
		S := scanner.Text()
		if !isNotEmpty(S) || !sameLength(num, N) {
			fmt.Println("指定した文字数の文字を入力してください。")
			goto L2
		}


		var count int = 0
		for i := 0; i < num; i++ {
			if N[i] != S[i] {
				count++
			}
		}

		fmt.Printf("%d\n", count)
		break
	}

}

func isNotEmpty(s string) bool {
	if s == "" {
		return false
	} else {
		return true
	}
}

func sameLength(i int, s string) bool {
	if i != len(s) {
		return false
	} else {
		return true
	}
}