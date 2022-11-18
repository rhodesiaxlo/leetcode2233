package main

import (
	"sort"
)

// interval 是一系列坐标， 坐标中 l, r 可能有重合，需要你合并重合的坐标， 输出一个新的坐标
func combineArr(intervals [][]int) [][]int {

	// 二位数组的初始化
	/*a := [][2]int {{1,2},{3,4}}
	fmt.Print(a)*/

	// sorting
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	sz := len(intervals)
	i:=0
	var ans [][]int
	for i=0;i < sz -1;i++ {
		if overlapping(intervals[i], intervals[i+1]) {
			intervals[i+1][0] = min(intervals[i][0], intervals[i+1][0])
			intervals[i+1][1] = max(intervals[i][1], intervals[i+1][1])
		} else {
			ans = append(ans, intervals[i])
		}
	}

	if i == sz - 1 {
		ans = append(ans, intervals[sz-1])
	}
	return ans
}

func max(i, j int) int{
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int{
	if i < j {
		return i
	}
	return j
}

func overlapping(l, r []int) bool {
	if l[0] < r[0] && l[1] >= r[0] {
		return true
	}

	return false
}


func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	i := 0
	for i < len(intervals) {
		left, right := intervals[i][0], intervals[i][1]
		j := i+1
		for j < len(intervals) && intervals[j][0] <= right {
			right = max(right, intervals[j][1])
			j++
		}
		res = append(res, []int{left, right})
		i = j
	}

	return res
}
