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
		scArray := strings.Split(scanner.Text(), " ")

		if len(scArray) != 4 {
			fmt.Printf("%vの入力を確認しました。整数を4つ入力してください\n", len(scArray))
			continue
		}

		var scNumberArray []int
		for _, v := range scArray {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("%v は整数ではないです。整数を入力してください。", v)
			} else if vInt < 0 {
				fmt.Printf("%v は正の整数ではないです。正の整数を入力してください。", v)
			} else {
				scNumberArray = append(scNumberArray, vInt)
			}
		}

		bigNumber := scNumberArray[0]
		for _, v := range scNumberArray {
			if bigNumber <= v {
				bigNumber = v
			}
		}

		fmt.Println(bigNumber)
		break
	}
}