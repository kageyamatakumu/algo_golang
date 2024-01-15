package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	texts := make([]string, 3)
	scanner := bufio.NewScanner(os.Stdin)

	for i := len(texts); i > 0 ; i-- {
		scanner.Scan()
		texts[i-1] = scanner.Text()
	}

	fmt.Println(strings.Join(texts, ""))
}