package heap

import (
	"leetcode/leetcode/sorting"
	"math"
	"testing"
)

func TestMaxHeap_Insert(t *testing.T) {
	// 第一种构造 maxheap 的方式
	nums := sorting.GenArray(1000000, 0)
	mh := NewMaxHeapByInsert(nums)

	// 排序的方式是每次 extraxMin
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i], _ = mh.ExtractMax()
	}
	if !sorting.IsSorted(nums) {
		t.Fatal("heap construct by insert, sort by extract error occurred")
	}
}

func TestMaxHeap_Heapify(t *testing.T) {
	nums := sorting.GenArray(1000000, 0)
	mh := NewMaxHeapByHeapify(nums)
	// 排序的方式是每次 extraxMin
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i], _ = mh.ExtractMax()
	}
	if !sorting.IsSorted(nums) {
		t.Fatal("heap construct by heapify, sort by extract error occurred")
	}
}

func TestMaxHeap_SortAsc(t *testing.T) {
	nums := sorting.GenArray(1000000, 0)
	mh := NewMaxHeapByHeapify(nums)
	mh.SortAsc()
	if !sorting.IsSorted(mh.GetArr()) {
		t.Fatal("heap construct by heapify, sort by extract error occurred")
	}
}

func TestMaxHeap_Insert2(t *testing.T) {
	// 第一种构造 maxheap 的方式
	nums := sorting.GenArray(100, 0)
	mh := NewMaxHeapByInsert(nums)

	// 排序的方式是每次 extraxMin
	mh.Insert(math.MaxInt32)
	m, _ := mh.ExtractMax()
	if m != math.MaxInt32 {
		t.Fatal("heap construct by insert, sort by extract error occurred")
	}
}

func TestMaxHeap_Size(t *testing.T) {
	nums := sorting.GenArray(100, 0)
	mh := NewMaxHeapByHeapify(nums)
	if mh.Size() != len(nums) {
		t.Fatal("max heap size func error occured")
	}
}

