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

		scanner.Scan()
		scNumber := scanner.Text()

		if num, err := strconv.Atoi(scNumber); err != nil {
			fmt.Println("整数を入力してください")
			continue
		} else if num < 0 {
			fmt.Println("正の整数を入力してください")
			continue
		} else {
			fmt.Println(string(scText[num-1]))
			break
		}
	}

}