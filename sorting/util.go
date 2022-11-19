package sorting

import (
	"math/rand"
	"time"
)

func GenArray(n, m int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < m; i++ {
		tmp := r.Intn(n - 1)
		arr[tmp], arr[n-1-tmp] = arr[n-1-tmp], arr[tmp]
	}
	return arr
}

func GenArrayDesc(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n-1-i
	}
	return arr
}

func IsSorted(nums []int) bool {
	sz := len(nums)
	if sz <= 2 {
		return true
	}

	for i := 0; i < sz-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}

	return true
}

func IsSortedDesc(nums []int) bool {
	sz := len(nums)
	if sz <= 2 {
		return true
	}

	for i := 0; i < sz-1; i++ {
		if nums[i] < nums[i+1] {
			return false
		}
	}

	return true
}

func GenArrWithSmallRange(n, m int) []int {
	nums := make([]int, n)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i:=0;i < n;i++ {
		nums[i] = r.Intn(m)
	}
	return nums
}
