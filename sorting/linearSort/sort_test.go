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

func TestHeapSortInsert(t *testing.T) {
	nums := sorting.GenArrayDesc(1000000)
	err := HeapSortInsert(nums)
	if err != nil {
		t.Fatal(" heap sort error occuring")
	}
	if !sorting.IsSorted(nums) {
		t.Fatal("heap sort not working")
	}
}

func TestHeapSortHeapify(t *testing.T) {
	nums := sorting.GenArrayDesc(1000000)
	err := HeapSortHeapify(nums)
	if err != nil {
		t.Fatal(" heap sort error occuring")
	}
	if !sorting.IsSorted(nums) {
		t.Fatal("heap sort not working")
	}
}

// heapSort1,heapsort2 都额外开辟了空间，实际上并不需要
func TestHeapSort(t *testing.T) {
	nums := sorting.GenArrayDesc(100)
	var err error
	nums, err = HeapSort(nums)
	if err != nil {
		t.Fatal(" heap sort error occuring")
	}
	if !sorting.IsSorted(nums) {
		t.Fatal("heap sort not working")
	}
}

func TestHeapSortDesc(t *testing.T) {
	nums := sorting.GenArrayDesc(1000000)
	nums, err := HeapSortDesc(nums)
	if err != nil {
		t.Fatal(" heap sort error occuring")
	}
	if !sorting.IsSortedDesc(nums) {
		t.Fatal("heap sort not working")
	}
}

func TestQuickSort2(t *testing.T) {
	nums := sorting.GenArrayDesc(1000000)
	QuickSort(nums)
	if !sorting.IsSorted(nums) {
		t.Fatal("quick sort not working")
	}
}
