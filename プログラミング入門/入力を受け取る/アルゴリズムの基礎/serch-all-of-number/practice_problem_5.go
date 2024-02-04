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

		if num < 1 {
			fmt.Println("1より小さい数値が入力されました。")
			continue
		}

    for i := 1; i <= num; i++ {
      if i % 5 == 0 && i % 3 == 0 {
        fmt.Println("FizzBuzz")
      } else if i % 3 == 0 {
        fmt.Println("Fizz")
      } else if i % 5 == 0 {
        fmt.Println("Buzz")
      } else {
        fmt.Println(i)
      }
    }

		break
	}

}