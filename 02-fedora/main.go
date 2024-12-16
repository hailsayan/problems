package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	nameCount := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		part := strings.Split(line, " ")
		firstName := part[0]
		nameCount[firstName]++
	}

	max := 0
	for _, count := range nameCount {
		if count > max {
			max = count
		}
	}
	fmt.Println(max)
}
