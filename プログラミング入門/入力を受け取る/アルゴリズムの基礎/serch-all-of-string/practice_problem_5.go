package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		S := scanner.Text()
		if !isNotEmpty(S) {
			fmt.Println("空文字が入力されました。")
			continue
		}

		L:
		scanner.Scan()
		T := scanner.Text()
		if !isNotEmpty(S) {
			fmt.Println("空文字が入力されました。")
			goto L
		}

		result := strings.Contains(S, T)

		if result {
			fmt.Println("Yes")
			break
		} else {
			fmt.Println("No")
			break
		}
	}


}

func isNotEmpty(s string) bool {
	if s == "" {
		return false
	} else {
		return true
	}
}