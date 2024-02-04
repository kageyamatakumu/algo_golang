package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func gcd(a, b int) int {
	if a == b {
		return a
	}

  if a < b {
    a, b = b, a
  }

	if a % b == 0 {
		return b
	}

	for i := b / 2; i > 0; i-- {
		if a % i == 0 && b % i == 0 {
			return i
		}
	}

	return 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		scTexts := scanner.Text()
		scTextsArray := strings.Fields(scTexts)

		var scNums []int
		for _, v := range scTextsArray {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("%vは正の整数ではありません。\n", v)
				continue
			}

			scNums = append(scNums, num)
		}

		if len(scNums) != 2 {
			fmt.Println("入力値が2つではありません。")
			continue
		}

		fmt.Printf("%d\n", gcd(scNums[0], scNums[1]))
		break
	}
}