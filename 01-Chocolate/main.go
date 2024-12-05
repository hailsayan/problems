package main

import (
	"fmt"
)

func main() {
	var n, x, k int
	fmt.Scan(&n, &x, &k)

	chocolate := x * (k - 1) * k / 2

	fmt.Println(chocolate)
}
