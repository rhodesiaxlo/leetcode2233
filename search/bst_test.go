package search

import (
	"leetcode/leetcode/sorting"
	"testing"
)

func TestBST(t *testing.T) {
	nums := sorting.GenArrayDesc(10000)
	index := BST(nums, 100)
	if index != 100 {
		t.Log("bst not working")
	}
}

func TestBSTFC(t *testing.T) {
	nums := sorting.GenArrayDesc(1000)
	nums[500] = 500
	f, c := BSTFC(nums, 500)
	if f != 499 || c != 500 {
		t.Fatal(" bst floor ceil not working")
	}
}

func TestFloor(t *testing.T) {
	nums := sorting.GenArray(100, 0)
	index := Floor(nums, 99)
	if index != 99 {
		t.Fatal("floor function accurate match not  working")
	}

	nums[99] = 100
	index2 := Floor(nums, 99)
	if index2 != 98 {
		t.Fatal("floor function fuzzey match not working")
	}

}

func TestCeil(t *testing.T) {
	nums := sorting.GenArray(100, 0)
	index := Ceil(nums, 99)
	if index != 99 {
		t.Fatal("ceil function accurate match not  working")
	}

	nums[98] = 97
	index2 := Ceil(nums, 98)
	if index2 != 99 {
		t.Fatal("ceil function fuzzy match not working")
	}
}
