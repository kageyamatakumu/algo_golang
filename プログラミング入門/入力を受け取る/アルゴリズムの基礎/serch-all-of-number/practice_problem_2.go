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

		var divisor int

		for i := 1; i <= num; i++ {
			if num % i == 0 {
				divisor++
			}
		}

		fmt.Printf("%d\n", divisor)
		break
	}
}