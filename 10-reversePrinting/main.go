package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var numbers []int

	for {
		line, _ := reader.ReadString('\n')
		line = line[:len(line)-1]
		number, _ := strconv.Atoi(line)

		if number == 0 {
			break
		}

		numbers = append(numbers, number)
	}

	for i := len(numbers) - 1; i >= 0; i-- {
		fmt.Println(numbers[i])
	}
}
