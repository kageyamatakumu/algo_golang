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

		var count int
		for i:= 1; i <= num; i++ {
			if i % 5 != 0 && i % 3 != 0 && i % 2 != 0 {
				count++;
			}
		}

		fmt.Printf("%d\n", count)
		break
	}

}