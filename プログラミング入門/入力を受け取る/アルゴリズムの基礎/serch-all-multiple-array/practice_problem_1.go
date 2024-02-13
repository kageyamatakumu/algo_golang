package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"errors"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		inputCountString, err := readInput(scanner, 2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		inputCountNums, err := convertToNumber(inputCountString)
		if err != nil {
			fmt.Println(err)
			continue
		}

		scanner.Scan()
		array1String, err := readInput(scanner, inputCountNums[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		array1Nums, err := convertToNumber(array1String)
		if err != nil {
			fmt.Println(err)
			continue
		}

		scanner.Scan()
		array2String, err := readInput(scanner, inputCountNums[1])
		if err != nil {
			fmt.Println(err)
			continue
		}

		array2Nums, err := convertToNumber(array2String)
		if err != nil {
			fmt.Println(err)
			continue
		}

		answer := greaterCount(array1Nums, array2Nums)

		fmt.Println(answer)
		break

	}

}

func readInput(scanner *bufio.Scanner, length int) ([]string, error) {
	scTexts := scanner.Text()
	scTextsArray := strings.Fields(scTexts)

	if len(scTextsArray) != length {
		return nil, errors.New("入力個数が一致しません")
	}

	return scTextsArray, nil
}

func convertToNumber(texts []string) ([]int, error) {
	var nums []int
	for _, v := range texts {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("数値ではないものが入力されています")
		}

		nums = append(nums, num)
	}

	return nums, nil
}

func greaterCount(array1 []int, array2 []int) (int) {

	var count int

	for i:= 0; i < len(array1); i++ {
		for j:= 0; j < len(array2); j++ {
			if array1[i] > array2[j] {
				count++
			}
		}
	}

	return count
}

