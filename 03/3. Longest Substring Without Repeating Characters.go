package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]int)
	maxLen := 0
	start := 0

	for end := 0; end < len(s); end++ {
		if index, found := charMap[s[end]]; found && index >= start {
			start = index + 1
		}

		charMap[s[end]] = end

		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
