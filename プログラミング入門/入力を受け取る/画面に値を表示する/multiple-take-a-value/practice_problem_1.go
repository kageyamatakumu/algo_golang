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
	var answer int

	for {
		scanner.Scan()
		scText := strings.Split(scanner.Text(), " ")

		for _, v := range scText {
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("整数を入力してください")
				continue
			}
			if num <= 0 {
				fmt.Println("正の整数を入力してください")
			}
			answer += num
		}
		break
	}

	if answer != 0 {
		fmt.Println(answer)
	}
}