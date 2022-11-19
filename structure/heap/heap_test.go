package heap

//
//func TestMaxHeap_Heapify(t *testing.T) {
//	nums := sorting.GenArray(10, 0)
//	mh := MaxHeap{}
//	mh.Heapify(nums)
//	linearSort.QuickSort(nums)
//	max := nums[len(nums)-1]
//	if mh.Max() != max {
//		t.Fatal("heapify not working")
//	}
//}
//
//func TestMaxHeap_ExtractMax(t *testing.T) {
//	nums := sorting.GenArray(10, 0)
//	mh := MaxHeap{}
//	mh.Heapify(nums)
//
//	sz := len(nums)
//	arr := make([]int, sz)
//
//	for i := sz - 1; i >= 0; i-- {
//		arr[i], _ = mh.ExtractMax()
//	}
//	if !sorting.IsSorted(arr) {
//		t.Fatal("max heap extract max is not working")
//	}
//}
//
//func TestMaxHeap_Insert(t *testing.T) {
//	nums := sorting.GenArray(10, 0)
//	mh := MaxHeap{}
//	mh.Heapify(nums)
//
//	mh.Insert(math.MaxInt32)
//	max := mh.Max()
//	if max != math.MaxInt32 {
//		t.Fatal("max heap insert not working")
//	}
//}
//
//func TestMinHeap_Heapify(t *testing.T) {
//	nums := sorting.GenArrayDesc(10)
//	mh := MinHeap{}
//	mh.Heapify(nums)
//
//	// lol , min 的 sort 不是这样子的
//	arr := make([]int, len(nums))
//	for !mh.Empty() {
//		tmpMin, _ := mh.ExtractMin()
//		arr = append(arr, tmpMin)
//	}
//	if !sorting.IsSorted(arr) {
//		t.Fatal("min heap not working")
//	}
//}
//
//func TestMinHeap_Insert(t *testing.T) {
//	nums := sorting.GenArrayDesc(10)
//	mh := MinHeap{}
//	mh.Heapify(nums)
//
//	mh.Insert(math.MinInt32)
//	min, _ := mh.ExtractMin()
//	if min!=math.MinInt32 {
//		t.Fatal("minheap extract insert not working")
//	}
//}
