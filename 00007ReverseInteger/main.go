package main

import (
	"fmt"
	"math"
)

func main() {
	i := 123
	fmt.Println(reverseInteger1(i))
}

/*
方法1
*/
func reverseInteger1(i int) int {
	symbol := 1
	if i < 0 {
		symbol = -1
	}
	div := symbol * i
	remaining := 0
	ret := 0
	for {
		if div == 0 {
			return ret * symbol
		}

		remaining = div % 10
		div = div / 10
		if math.MaxInt32 - 10 * ret < remaining {
			// overflow
			return 0
		}
		ret = 10 * ret + remaining
	}
}
