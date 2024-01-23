package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		scText := scanner.Text()
		num, err := strconv.Atoi(scText)
		if err != nil {
			fmt.Printf("%vは正の整数ではありません。\n", scText)
			continue
		}
		if num < 0 {
			fmt.Printf("%vは正の整数ではありません。\n", scText)
			continue
		}

		var leftCount, rightCount int

		L:
		for i := 0; i < num; i++ {
			scanner.Scan()
      switch scanner.Text() {
      case "left":
        leftCount++
      case "right":
        rightCount++
      default:
        fmt.Println("[left]か[right]を入力してください。")
        goto L
      }
		}

    switch {
    case leftCount > rightCount:
      fmt.Println("left")
    case leftCount < rightCount:
      fmt.Println("right")
    default:
      fmt.Println("same")
    }

		break
	}

}