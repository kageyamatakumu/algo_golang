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
		firstInputString, err := readInput(scanner, 3)
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

		L2:
		scanner.Scan()
		thirdInputString, err := readInput(scanner, firstInputNums[1])
		if err != nil {
			fmt.Println(thirdInputString)
			goto L2
		}

		thirdInputNums, err := convertToNumber(thirdInputString)
		if err != nil {
			fmt.Println(err)
			goto L2
		}

		L3:
		scanner.Scan()
		forthInputString, err := readInput(scanner, firstInputNums[2])
		if err != nil {
			fmt.Println(err)
			goto L3
		}

		forthInputNums, err := convertToNumber(forthInputString)
		if err != nil {
			fmt.Println(err)
			goto L3
		}

		answer := count_matching_triplets(secondInputNums, thirdInputNums, forthInputNums)
		fmt.Println(answer)
		break
	}
}

func readInput(scanner *bufio.Scanner, count int) ([]string, error) {
	scTexts := scanner.Text()
	scTextsArray := strings.Fields(scTexts)

	if len(scTextsArray) != count {
		return nil, errors.New(strconv.Itoa(count) + "つの数値を入力してください")
	}

	return scTextsArray, nil
}

func convertToNumber(array []string)([]int, error) {
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

func count_matching_triplets(aNums []int, bNums []int, cNum []int) int {
	var count int

	for i := 0; i < len(aNums); i++ {
		for j := 0; j < len(bNums); j++ {
			for k := 0; k < len(cNum); k++ {
				if aNums[i] + bNums[j] == cNum[k] {
					count++
				}
			}
		}
	}

	return count
}