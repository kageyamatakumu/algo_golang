package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		inputString, err := readInput(scanner, 2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		inputNums, err := coverToNumber(inputString)
		if err != nil {
			fmt.Println(err)
			continue
		}

		answer, err := count_pairs_same_ones_digit_in_range(inputNums[0], inputNums[1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(answer)
		break
	}

}

func readInput(scanner *bufio.Scanner, length int) ([]string, error) {
	scTexts := scanner.Text()
	scTextsArray := strings.Fields(scTexts)

	if len(scTextsArray) != length {
		return nil, errors.New(strconv.Itoa(length) + "つの数値を入力してください")
	}

	return scTextsArray, nil
}

func coverToNumber(array []string) ([]int, error) {
	var nums []int
	for _, v := range array {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("数値ではないのものが入力されています")
		}

		nums = append(nums, num)
	}

	return nums, nil
}

func count_pairs_same_ones_digit_in_range(aNum int, bNum int) (int, error) {
	if aNum > bNum {
    return 0, errors.New("1つ目の数字は2つ目の数字以下である必要があります")
  }
	var count int

	for i := aNum; i < bNum; i++ {
		for j := i+1; j <= bNum; j++ {
			if i % 10 == j % 10 {
				count++
			}
		}
	}

	return count, nil
}