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
		firstInputString, err := readInput(scanner, 2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		firstInputNums, err := coverToNumber(firstInputString)
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

		secondInputNums, err := coverToNumber(secondInputString)
		if err != nil {
			fmt.Println(err)
			goto L
		}

		answer := count_pairs_sum_under(firstInputNums, secondInputNums)
		fmt.Println(answer)
		break
	}

}

func readInput(scanner *bufio.Scanner, length int) ([]string, error) {
	scTexts := scanner.Text()
	scTextsArray := strings.Fields(scTexts)

	if len(scTextsArray) != length {
		return nil, errors.New(strconv.Itoa(length) + "つの整数を入力してください")
	}

	return scTextsArray, nil
}

func coverToNumber(array []string) ([]int, error) {
	var nums []int

	for _, v := range array {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("数位ではないものが入力されています")
		}

		nums = append(nums, num)
	}

	return nums, nil
}

func count_pairs_sum_under(aNums []int, bNums []int) int {
	var count int

	for i := 0; i < aNums[0]; i++ {
		for j := i+1; j < aNums[0]; j++ {
			if bNums[i] + bNums[j] <= aNums[1] {
				count++
			}
		}
	}

	return count
}