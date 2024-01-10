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
		X, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("整数を入力してください")
			continue
		}

		if X < 0 || X > 23 {
			fmt.Println("0以上23以下")
			continue
		}

		fmt.Printf("%v\n", 24 - X)
		break
	}
}