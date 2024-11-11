package main

import (
	"fmt"
)

func sum(x, z int) int {
	return x + z
}

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println(sum(a, b))
}
