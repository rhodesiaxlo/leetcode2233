package main

import (
	"fmt"
)

func main() {
	//i := 12321
	j := 10
	fmt.Println(isPalindrom2(j))
}

/*
暴力解法， 求出 int 的 rever 然后比较 不考虑溢出的情况
*/
func isPalindrom1(i int) bool {
	if i < 0 {
		return false
	}

	rev := reverse(i)
	return i == rev
}

/*
上下位  ts: 99%, ss: 100%
*/
func isPalindrom2(x int) bool {
	if x < 0 {
		return false
	}
	// get range
	n, j := 1, x
	for j/10 > 0 {
		n *= 10
		j = j / 10
	}

	i := x
	for n >= 10 {
		low := i % 10
		hi := i / n
		if hi != low {
			return false
		}

		i = (i - hi*n) / 10
		// range reduced by 100
		n /= 100
	}

	return true
}

/*
	根据 panlindrom 的性质， 将数字拆分位长度相等的2部分， 然后再比较大小
	ts:34%  ss:64%
*/
func isPalindrom3(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}

	rev := 0
	for x > rev {
		rev = 10 * rev + x % 10
		x /= 10
	}

	return rev == x || x == rev/10
}

func reverse(i int) int {
	front := i
	back := 0
	for front > 0 {
		remain := front % 10 // get mod first
		front = front / 10
		back = 10*back + remain
	}
	return back
}
