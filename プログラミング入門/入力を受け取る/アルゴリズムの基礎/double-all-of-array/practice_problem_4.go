package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		text, err := readInput(scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var charArray []string
		for _, v := range text {
			if contains(charArray, string(v)) {
				charArray = append(charArray, string(v))
			}
		}

		fmt.Println(len(charArray))
		break
	}
}

func readInput(scanner *bufio.Scanner) (string, error) {
	scText := scanner.Text()

	for _, v := range scText  {
		if rune(v) < 97 || rune(v) > 122 {
			return "", fmt.Errorf("小文字以外が入力されました")
		}
	}

	return scText, nil
}

func contains(elems []string, s string) bool {
	for _, v := range elems {
		if s == v {
			return false
		}
	}

	return true
}