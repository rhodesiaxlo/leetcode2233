package linearSort

import (
	"leetcode/leetcode/structure/heap"
)

// 排序的几个维度
// 1. stable // 可以多次拍， 比如成绩倒序， 然后名字的字典序
// 2. 时间复杂度  最好， 最差， 平均
// 3. inplace
// 4. 是不是一定要先有序 heap
// 5. 额外的空间 和 3 有点类似， 但是是定量分析

/*
升序排列
*/
func BubbleSort(nums []int) {
	sz := len(nums)
	for i := 0; i < sz-1; i++ {
		for j := 0; j < sz-i-1; j++ { // 还要判定 j+1, 所以 j < sz -i - 1
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 为什么选择排序是不稳定的， 因为如果后面有更小的话， 会掉个
func SelectionSort(nums []int) {
	sz := len(nums)
	for i := 0; i < sz-1; i++ {
		minIndex := i
		for j := i + 1; j < sz; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
}

/**
插入排序比选择排序改进的点在于， 在近乎有序的数组中接近O(n)
在无需的数组中也比选择排序好
*/
func InsertSort(nums []int) {
	l, r := 0, len(nums)-1
	__insertSort(nums, l, r)
}

func __insertSort(nums []int, l, r int) {
	for i := l + 1; i <= r; i++ {
		for j := i; j > l && nums[j] < nums[j-1]; j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}

func __insertSort2(nums []int, l, r int) {
	for i := l + 1; i <= r; i++ {
		tmp := nums[i]
		j := i
		for j = i; j > l && nums[j] < tmp; j-- {
			nums[j] = nums[j-1]
		}
		// 找到第一个
		nums[j] = tmp
	}
}

/**
归并排序
*/
func MergeSort(nums []int) {
	__mergeSort(nums, 0, len(nums)-1)
}

/**
nums 待排序数组， l, 左边界  m 中间点 r 右边界
*/
func __mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}

	//if r-l <= 15 {
	//	InsertSort2(nums, l, r)
	//	return
	//}

	mid := (r-l)/2 + l
	__mergeSort(nums, l, mid)
	__mergeSort(nums, mid+1, r)
	__merge(nums, l, mid, r)

	//if nums[mid] > nums[mid+1] {
	//	__merge(nums, l, mid, r)
	//}
}

func __merge(nums []int, l, m, r int) {
	aux := make([]int, r-l+1)

	for i := l; i <= r; i++ {
		aux[i-l] = nums[i]
	}

	i, j, k := l, m+1, l
	// use aux to sort
	for {
		if i > m && j > r {
			break
		}
		if i > m {
			nums[k] = aux[j-l]
			j++
		} else if j > r {
			nums[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			nums[k] = aux[i-l]
			i++
		} else {
			nums[k] = aux[j-l]
			j++
		}
		k++
	}
}

/**

 */
func MergeSortBU() {
	// todo...
}

func QuickSortLLPtr() {
	// todo...
}

func QuickSortLRPtr() {
	// todo..
}

// 如何理解快排是一种不稳定的排序方式
func QuickSort(nums []int) {
	__quickSort(nums, 0, len(nums)-1)
}

func __quickSort(nums []int, l, r int) {
	if r-l <= 15 {
		__insertSort2(nums, l, r)
		return
	}

	// partition
	p := __partition(nums, l, r)

	// quick sort
	__quickSort(nums, l, p-1)
	__quickSort(nums, p+1, r)
}

func __partition(nums []int, l, r int) int {
	j := l
	tmp := nums[l]
	// 约定 (-infi, j) 的元素都小于 tmp
	// 遍历数组, 如果发现有比 tmp 下的 放到 j 的位置 j++
	// 随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	//swap( nums[l] , nums[rand()%(r-l+1)+l] );
	for i := l + 1; i <= r; i++ {
		if nums[i] < tmp {
			nums[j] = nums[i]
			j++
		}
	}

	nums[j] = tmp

	return j
}

/**
有一个缺点就是如果有大量重复元素的话，会浪费性能所有产生了三路快排
*/
func QuickSortThreeWays() {
	// todo..
}

// a lot of same values
func CountingSort(nums []int) {

}

func __countingSort(nums []int, biggest, smallest int) {

}

func RadixSortLSD(nums []int) {
	// todo...
}

func RadixSortMSD(nums []int) {
	// todo..
}

// val is dense dis
// len(n) ≈ largest - smallest
func PigeonholdSort(nums []int) {
	// todo.
}

func BucketSort(nums []int) {
	// todo..
}

func StdSort(nums []int) {
	// todo
}

func StdStableSort(nums []int) {
	// todo
}

func ShellSort(nums []int) {
	// todo
}

func CocktailShakerSort(nums []int) {
	// todo..
}

func GnomeSort(nums []int) {
	// todo
}

func BitonicSort(nums []int) {
	// todo...
}

func BogoSort(nums []int) {

}

// cocktail shaker sort && optimized bubble sort
func ExchangeSort() {
	// todo..
}

func HeapSortHeapify(nums []int) error {
	mh := heap.MinHeap{}
	mh.Heapify(nums)
	for i := 0; i < len(nums); i++ {
		var err error
		nums[i], err = mh.ExtractMin()
		if err != nil {
			return err
		}
	}
	return nil
}

// 可以看到，不管是 heapSort1 还是 heapSort2， 都额外开辟了 n 个元素的空间, 实际上
func HeapSortInsert(nums []int) error {
	mh := heap.MinHeap{}
	sz := len(nums)
	for i := 0; i < sz; i++ {
		mh.Insert(nums[i])
	}

	for i := 0; i < sz; i++ {
		var err error
		nums[i], err = mh.ExtractMin()
		if err != nil {
			return err
		}
	}
	return nil
}

func  HeapSort(nums []int) error {
	mh := heap.MaxHeap{}
	mh.Heapify(nums)
	mh.SortAsc()
	nums = mh.GetArr()
	return nil
}

func  HeapSortDesc(nums []int) error {
	mh := heap.MinHeap{}
	mh.Heapify(nums)
	mh.SortDesc()
	nums = mh.GetArr()
	return nil
}

