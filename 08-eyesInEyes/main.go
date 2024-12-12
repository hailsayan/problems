package main

import (
	"fmt"
)

func main() {
	var row1, row2 [8]int
	for i := 0; i < 8; i++ {
		fmt.Scan(&row1[i])
	}
	for i := 0; i < 8; i++ {
		fmt.Scan(&row2[i])
	}
	count := 0
	for i := 0; i < 8; i++ {
		if row1[i] == 1 && row2[i] == 1 {
			count++
		}
	}
	fmt.Println(count)
}
