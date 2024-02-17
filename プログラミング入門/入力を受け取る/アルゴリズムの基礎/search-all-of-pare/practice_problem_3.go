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
		firstInputString, err := readInput(scanner, 1)
		if err != nil {
			fmt.Println(err)
			continue
		}

		firstInputNums, err := convertToNumber(firstInputString)
		if err != nil {
			fmt.Println(err)
			continue
		}

		L:
		scanner.Scan()
		secondInputString, err := readInput(scanner, firstInputNums[0])
		if err != nil {
			fmt.Println(err)
			goto L
		}

		secondInputNums, err := convertToNumber(secondInputString)
		if err != nil {
			fmt.Println(err)
			goto L
		}

		answer := count_max_ai_equal_aj_in_ak(secondInputNums)
		fmt.Println(answer)
		break
	}
}

func readInput(scanner *bufio.Scanner, length int) ([]string, error) {
	scTexts := scanner.Text()
	scTextsArray := strings.Fields(scTexts)

	if len(scTextsArray) != length {
		return nil, errors.New(strconv.Itoa(length)+"つの数値を入力してください")
	}

	return scTextsArray, nil
}

func convertToNumber(array []string) ([]int, error) {
	var nums []int

	for _, v := range array {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("数値ではないものが入力されています")
		}

		nums = append(nums, num)
	}

	return nums, nil
}

func count_max_ai_equal_aj_in_ak(array []int) int {
	var count int

	for i := 0; i < len(array); i++ {
		for j := i+1; j < len(array); j++ {
			for k := j+1; k < len(array); k++ {
				if array[j] >= array[i] && array[j] >= array[k] {
					count++
				}
			}
		}
	}

	return count
}