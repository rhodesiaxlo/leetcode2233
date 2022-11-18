package linearSort

import (
	"fmt"
	"leetcode/leetcode/sorting"
	"testing"
)

// 100k 需要 50 多秒
func TestBubbleSort(t *testing.T) {
	nums := [][]int{}
	for i := 0; i < 10; i++ {
		//nums = append(nums, sorting.GenArray(10000, 100))
		//nums = append(nums, sorting.GenArray(100000, 1000))
	}

	for _, v := range nums {
		BubbleSort(v)
		if !sorting.IsSorted(v) {
			t.Fatalf("%s not sorting", fmt.Sprint(v))
		}
	}
}

func TestInsertSort0(t *testing.T) {
	nums := sorting.GenArray(1000000, 1000)
	InsertSort(nums)
	if !sorting.IsSorted(nums) {
		t.Fatal("nearlly sorted array")
	}
}

func TestInsertSort1(t *testing.T) {
	nums := sorting.GenArray(1000000, 1)
	InsertSort(nums)
	if !sorting.IsSorted(nums) {
		t.Fatal("nearlly sorted array")
	}
}

func TestMergeSort(t *testing.T) {
	nums := sorting.GenArrayDesc(1000000)
	MergeSort(nums)
	if !sorting.IsSorted(nums) {
		t.Fatal("nearlly sorted array")
	}
}

func TestQuickSort(t *testing.T) {
	nums := sorting.GenArrayDesc(1000000)
	MergeSort(nums)
	if !sorting.IsSorted(nums) {
		t.Fatal("nearlly sorted array")
	}
}
