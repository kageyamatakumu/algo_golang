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
		scText := strings.Split(scanner.Text(), " ")

		num1, err := strconv.Atoi(scText[0])
    if err != nil {
      fmt.Println("整数を入力してください")
      continue
    }
    num2, err := strconv.Atoi(scText[1])
    if err != nil {
      fmt.Println("整数を入力してください")
      continue
    }

    if num1 % num2 == 0 {
      fmt.Println("Yes")
    } else {
      fmt.Println("No")
    }
    break

		// for i, v := range scText {
		// 	num, err := strconv.Atoi(v)
		// 	if err != nil {
		// 		fmt.Println("整数を入力してください")
		// 		continue
		// 	}
		// 	if num <= 0 {
		// 		fmt.Println("正の整数を入力してください")
		// 		continue
		// 	}
		// 	if i != 0 {
		// 		if value % num == 0 {
		// 			fmt.Println("Yes")
		// 		} else {
		// 			fmt.Println("No")
		// 		}
		// 	} else {
		// 		value = num
		// 	}
		// }
		// break
	}
}