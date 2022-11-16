package main

import (
	"fmt"
)

func main() {
	s := "III"
	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	m := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
	sz := len(s)
	ans := 0
	ans += m[s[sz-1:]]
	for i := sz - 2; i >= 0; i-- {
		if m[s[i:i+1]] < m[s[i+1:i+2]] {
			ans -= m[s[i:i+1]]
		} else {
			ans += m[s[i:i+1]]
		}
	}
	return ans
}
