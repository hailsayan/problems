package main

import "fmt"

func pigeonsPrinciple(n, m int) string {
	if n > m {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	fmt.Println(pigeonsPrinciple(n, m))
}
