package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		firstInput := readInput(scanner)
		firstInputNum, err := covertToNumber(firstInput)
		if err != nil {
			fmt.Println(err)
			continue
		}

		array := loopReadInput(scanner, firstInputNum)

		answer := sameText(array)
		fmt.Println(answer)
		break
	}
}

func readInput(scanner *bufio.Scanner) string {
	scText := scanner.Text()
	return scText
}

func covertToNumber(text string) (int, error) {
	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, errors.New("数値ではないものが入力されています")
	}

	return num, nil
}

func loopReadInput(scanner *bufio.Scanner, count int) ([]string) {
	var inputArray []string
	for i := 0; i < count; i++ {
		scanner.Scan()
		scText := readInput(scanner)

		inputArray = append(inputArray, scText)
	}

	return inputArray
}

func sameText(array []string) string {
	for i := 0; i < len(array); i++ {
		for j := i+1; j < len(array); j++ {
			if array[i] == array[j] {
				return "Yes"
			}
		}
	}

	return "No"
}