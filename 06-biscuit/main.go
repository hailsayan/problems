package main

import (
	"fmt"
)

func main() {
	var n, x, k int
	fmt.Scan(&n, &x, &k)

	totalChocolate := (k - 1) * x

	fmt.Println(totalChocolate)
}
