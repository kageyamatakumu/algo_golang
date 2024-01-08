package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(i % 5)
}