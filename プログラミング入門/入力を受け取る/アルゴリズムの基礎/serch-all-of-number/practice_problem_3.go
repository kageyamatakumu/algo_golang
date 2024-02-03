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
		} else if num < 2 {
			fmt.Printf("No")
			break
		}

		if num % 2 != 0 && num % 3 != 0 && num % 5 != 0 && num % 7 != 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

		break
	}
}