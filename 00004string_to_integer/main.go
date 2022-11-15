package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "   -4193 with words "
	fmt.Println(atoi(s))
}

func atoi(s string) int {
	s = strings.Trim(s, " ")
	var l int

	l = 0
	ret := 0
	c := s[0]
	if c == '+' || c == '-' {
		l++
	}

	// l is the start point
	for i := l; i < len(s); i++ {
		// read until we meet a non-digit
		if s[i] >= '0' && s[i] <= '9' {
			single := (int)(s[i] - '0')
			ret = ret * 10

			// how to clamp

			ret += single
			continue
		}
		break
	}

	if c == '-' {
		return -1 * ret
	}
	return ret
}
