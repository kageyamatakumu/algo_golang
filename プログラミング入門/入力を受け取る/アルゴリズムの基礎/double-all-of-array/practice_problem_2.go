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

		N, K, err := readInput(scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}

		L:
		answerCount, err := countDivisors(N, K)
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
			return 0, 0, fmt.Errorf("%vは正の整数ではありません", v)
		}

		scNums = append(scNums, num)
	}

	if scNums[0] <= scNums[1] {
    return 0, 0, fmt.Errorf("%dは%d以下ではありません", scNums[0], scNums[1])
  }

	return scNums[0], scNums[1], nil
}

func countDivisors(n, k int) (int, error) {
	if n <= 0 || k <= 0 {
    return 0, fmt.Errorf("nとkは正の整数である必要があります")
  }

	var answerCount int
	for i := 1; i <= n; i++ {
		count := 0
		for j := 1; j <= n; j++ {
			if i % j == 0  {
				count++
			}
		}

		if count == k {
			answerCount++
		}
	}

	return answerCount, nil
}
