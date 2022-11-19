package search

import "leetcode/leetcode/sorting/linearSort"

func BST(nums []int, x int) int {
	linearSort.QuickSort(nums)

	sz := len(nums)
	l, r := 0, sz-1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == x {
			return mid
		} else if nums[mid] > x {
			r = mid - 1
		} else if nums[mid] < x {
			l = mid + 1
		}
	}

	return -1
}

// 对于有重复元素的查询
func BSTFC(nums []int, x int) (floor, ceil int) {
	index := BST(nums, x)
	if index == -1 {
		return -1, -1
	}

	floor, ceil = index, index
	for nums[floor] == x {
		floor--
	}

	for nums[ceil] == x {
		ceil++
	}

	return floor+1, ceil-1
}

/**
floor 找出 x 在 nums 中的第一次出现的索引值，如果没有的话返回比 x 小的最大值在 nums 中第一次出现的索引
 */
func Floor(nums []int, x int) int {
	// todo
	// 这种模糊的搜索很有意思
	linearSort.QuickSort(nums)

	sz := len(nums)
	l, r := -1,sz-1
	for l < r {
		mid := (r-l+1)/2 + l
		if nums[mid] >= x { // 大于等于都继续向下找
			r =mid - 1
		} else {
			l = mid
		}
	}

	if l == r && nums[l+1] == x {
		return l+1
	}
	return l
}

func Ceil(nums []int, x int) int {
	linearSort.QuickSort(nums)

	sz := len(nums)
	l, r := 0, sz
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] <= x {
			l = mid +1
		} else {
			r = mid
		}
	}

	if l== r && nums[r-1] == x {
		return r-1
	}
	return r
}
