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

		L, R, err := readInput(scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}

		L:
		answerCount, err := countPalindrome(L, R)
		if err != nil {
			fmt.Println(err)
			goto L
		}

		fmt.Println(answerCount)
		break
	}
}

func readInput(scanner *bufio.Scanner) (int, int, error) {
	scTexts := scanner.Text()
	scTextsArray := strings.Fields(scTexts)

	var scNums []int
	for _, v := range scTextsArray {
		num, err := strconv.Atoi(v)
		if err != nil {
			return 0, 0, fmt.Errorf("%vは正の整数ではありません\n", v)
		}

		scNums = append(scNums, num)
	}

	if scNums[0] > scNums[1] {
    return 0, 0, fmt.Errorf("%dは%d以下ではありません", scNums[0], scNums[1])
  }

	return scNums[0], scNums[1], nil
}


func countPalindrome(l, r int) (int, error) {
	if l > r {
		return 0, fmt.Errorf("2回目に入力した値は最初の数値より大きい必要があります")
	}
	answerCount := 0
	for i := l; i <= r; i++ {
    if isPalindrome(i) {
      answerCount++
    }
	}

	return answerCount, nil
}

func isPalindrome(n int) bool {
  if n < 0 {
    return false
  }

  if n < 10 {
    return true
  }

  nStr := strconv.Itoa(n)

  length := len(nStr)

  for i := 0; i < length/2; i++ {
    if nStr[i] != nStr[length-1-i] {
      return false
    }
  }

  return true
}
