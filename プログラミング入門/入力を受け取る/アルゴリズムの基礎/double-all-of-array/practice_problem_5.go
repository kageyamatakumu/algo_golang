package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for{
		scanner.Scan()

		num, err := readInputNum(scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var textArray []string
		addTextArray(scanner, &textArray, num)

		var count int
		for _, v := range textArray {
			if IsPalindrome(v) {
				count++
			}
		}

		fmt.Println(count)
		break
	}
}

func readInputNum(scanner *bufio.Scanner) (int, error) {
	scText := scanner.Text()

	num, err := strconv.Atoi(scText)
	if err != nil {
		return 0, fmt.Errorf("数値以外が入力されました")
	}

	return num, nil
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

func addTextArray(scanner *bufio.Scanner, array *[]string, num int) {
  for i := 0; i < num; i++ {
    scanner.Scan()
    text, err := readInput(scanner)
    if err != nil {
      fmt.Println(err)
      continue
    }

    *array = append(*array, text)
  }
}

func IsPalindrome(text string) bool {
	for i := 0; i < len(text); i++ {
		if text[i] != text[len(text)-1-i] {
			return false
		}
	}

	return true
}