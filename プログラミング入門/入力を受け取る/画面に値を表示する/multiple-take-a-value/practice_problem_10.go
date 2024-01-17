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

		L:
			scanner.Scan()
			scTextNum := scanner.Text()
			scTextArray := strings.Split(scTextNum, " ")
			if len(scTextArray) != 2 {
				fmt.Println("空白区切りで2つ数値を入力してください")
				continue
			}

			var scNumberArray []int

			for _, v := range scTextArray {
				if num, err := strconv.Atoi(v); err != nil {
					fmt.Printf("%v: は数値ではありません。数値を入力してください\n", v)
					goto L
				} else {
					scNumberArray = append(scNumberArray, num)
				}
			}

			scTextRune := []rune(scText)

			box := scTextRune[scNumberArray[0] - 1]
			scTextRune[scNumberArray[0] - 1] = scTextRune[scNumberArray[1] - 1]
			scTextRune[scNumberArray[1] - 1] = box

			fmt.Println(string(scTextRune))
			break
		}
}