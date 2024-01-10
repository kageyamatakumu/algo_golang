package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var s string;

	for {
		scanner.Scan()
		if len(scanner.Text()) == 5 {
			s = scanner.Text()
			break
		} else {
			fmt.Println("長さ5の文字列を入力してください")
		}
	}

	fmt.Println(s[2:3])
}