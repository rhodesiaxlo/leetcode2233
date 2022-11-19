package heap

import (
	"leetcode/leetcode/sorting"
	"math"
	"testing"
)

func TestMinHeap_Insert(t *testing.T) {
	// 第一种构造 maxheap 的方式
	nums := sorting.GenArrayDesc(1000000)
	mh := NewMinHeapByInsert(nums)

	// 排序的方式是每次 extraxMin
	for i := 0; i < len(nums); i++ {
		nums[i], _ = mh.ExtractMin()
	}
	if !sorting.IsSorted(nums) {
		t.Fatal("min heap construct by insert, sort by extract error occurred")
	}
}

func TestMinHeap_Heapify(t *testing.T) {
	nums := sorting.GenArrayDesc(1000000)
	mh := NewMinHeapByHeapify(nums)
	// 排序的方式是每次 extraxMin
	for i := 0; i < len(nums); i++ {
		nums[i], _ = mh.ExtractMin()
	}
	if !sorting.IsSorted(nums) {
		t.Fatal("min heap construct by heapify, sort by extract error occurred")
	}
}

func TestMinHeap_SortDesc(t *testing.T) {
	nums := sorting.GenArrayDesc(1000000)
	mh := NewMinHeapByHeapify(nums)
	mh.SortDesc()
	if !sorting.IsSortedDesc(mh.GetArr()) {
		t.Fatal("min heap construct by heapify, sort by extract error occurred")
	}
}

func TestMinHeap_Insert2(t *testing.T) {
	// 第一种构造 maxheap 的方式
	nums := sorting.GenArrayDesc(100)
	mh := NewMinHeapByHeapify(nums)

	// 排序的方式是每次 extraxMin
	mh.Insert(math.MinInt32)
	m, _ := mh.ExtractMin()
	if m != math.MinInt32 {
		t.Fatal("min heap construct by insert, sort by extract error occurred")
	}
}

func TestMinHeap_Size(t *testing.T) {
	nums := sorting.GenArrayDesc(100)
	mh := NewMinHeapByHeapify(nums)
	if mh.Size() != len(nums) {
		t.Fatal("min heap size func error occured")
	}
}
