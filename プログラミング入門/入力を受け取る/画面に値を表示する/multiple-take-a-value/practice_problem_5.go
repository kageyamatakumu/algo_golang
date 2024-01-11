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
		scText := scanner.Text()
		slice := strings.Split(scanner.Text(), " ")

		if len(slice) != 3 {
			fmt.Printf("3つの整数を入力してください。%v個しか入力を受けていません。", len(scText))
		}

		var average int
		for _, v := range slice {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("整数以外の値がありました。整数を入力してください")
			} else {
				average += vInt
			}
		}

		fmt.Println(int(average / len(slice)))
		break
	}
}