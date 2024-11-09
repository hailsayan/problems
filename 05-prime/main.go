package main

import "fmt"

func primeEtc(n int) int {
	sumDigit := 0
	for n != 0 {
		digit := n % 10
		sumDigit += digit
		n /= 10
	}
	return sumDigit
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(primeEtc(n))
}
