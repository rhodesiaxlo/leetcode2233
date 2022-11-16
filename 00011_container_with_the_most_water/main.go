package main

import "fmt"

func main() {
	h := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(h))
}

func maxArea(height []int) int {
	ma, sz := 0, len(height)
	l, r := 0, sz-1
	for l < r {
		tmpArea := 0
		if height[l] > height[r] {
			tmpArea = (r - l) * height[r]
			r--
		} else {
			tmpArea = (r - l) * height[l]
			l++
		}
		ma = getMax(ma, tmpArea)
	}
	return ma
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
