package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	fibSet := make(map[int]bool)
	a, b := 1, 2
	fibSet[a] = true
	fibSet[b] = true
	for b <= n {
		// fmt.Println(a, b, fibSet)
		a, b = b, a+b
		fibSet[b] = true
	}

	result := ""
	for i := 1; i <= n; i++ {
		// fmt.Println(fibSet[i])
		if fibSet[i] {
			result += "+"
		} else {
			result += "-"
		}
	}
	fmt.Println(result)
}
