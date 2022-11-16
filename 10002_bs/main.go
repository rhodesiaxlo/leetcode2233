package main

import "fmt"
/**
1. l <= r
2. mid = (r-l)/2 + l
3. bs 有局限性，只能搜索不重复的数组
 */
func main() {
	src := []int {1,2}
	fmt.Println(bs(2, src))
}

/**
BS 搜索
前提数组已经排序
算法复杂度 log(n) 级别
空间复杂度 O(1)
*/
func bs(t int, src []int) int {
	if len(src) == 0 {
		return -1
	}

	sz := len(src)
	l, r := 0, sz - 1
	for l <= r { // 一定要有等号
		mid := (r-l)/2 + l
		if src[mid] > t {
			r = mid -1
		} else if src[mid] == t {
			return mid
		} else {
			l = mid + 1
		}
	}
	return -1
}
