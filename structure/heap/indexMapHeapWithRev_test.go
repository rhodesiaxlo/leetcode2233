package heap

import (
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// 测试insert
func TestIndexMaxHeapWithRev_Insert(t *testing.T) {
	n:=1000000
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}

	imh := IndexMaxHeapWithRev{}
	for i:=0;i<n;i++ {
		imh.Insert(Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}

	imh.SortAsc()

	ans := imh.GetArr()
	sz2 := len(ans)
	for i:=0;i < sz2-1;i++ {
		if ans[i].Gt(ans[i+1]) {
			t.Fatal("IndexMapHeapWithRev insert sort not working")
		}
	}
}

// heapify
func TestIndexMaxHeapWithRev_Heapify(t *testing.T) {
	n:=100
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}
	imh := IndexMaxHeapWithRev{}
	imh.Heapify(items)
	imh.SortAsc()
	ans := imh.GetArr()
	sz2 := len(ans)
	for i:=0;i < sz2-1;i++ {
		if ans[i].Gt(ans[i+1]) {
			t.Fatal("IndexMapHeapWithRev heapify sort not working")
		}
	}
}

func TestIndexMaxHeapWithRev_Max(t *testing.T) {
	n:=1000000
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}
	imh := IndexMaxHeapWithRev{}
	imh.Heapify(items)
	if imh.Max().v != n-1 {
		t.Fatal("IndexMaxHeapWithRev max function not working")
	}
}

// insert
func TestIndexMaxHeapWithrev_InsertAndSort(t *testing.T) {
	n:=1000000
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}
	imh := IndexMaxHeapWithRev{}
	imh.Heapify(items)
	imh.Insert(Item{
		v:    n,
		name: strconv.Itoa(n),
	})
	if imh.Max().v != n {
		t.Fatal("IndexMapHeapWithRev insert after Heapify not working")
	}
}

// change
func TestIndexMaxHeapWithRev_Change(t *testing.T) {
	n:=1000000
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}
	imh := IndexMaxHeapWithRev{}
	imh.Heapify(items)
	imh.Change(0, Item{math.MaxInt32, strconv.Itoa(math.MaxInt32)})
	if imh.Max().v != math.MaxInt32 {
		t.Fatal("IndexMapHeapWithRev change function not working")
	}
}

// getItems
func TestIndexMaxHeapWithRev_GetItem(t *testing.T) {
	n:=1000000
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}
	imh := IndexMaxHeapWithRev{}
	tmpItem, _:= imh.GetItem(0)
	if tmpItem.v != 0 {
		t.Fatal("IndexMapHeapWithRev getItem not working")
	}

}

func TestIndexMaxHeapWithRev_ChangeSoManyTimes(t *testing.T) {
	n:=1000000
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}

	imh := IndexMaxHeapWithRev{}
	imh.Heapify(items)
	m := 10000
	// 修改 10000 次，观察性能
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i:=0;i < m ;i++{
		tmpIndex := r.Intn(n)
		tmpV, _ := imh.GetItem(tmpIndex)
		tmpV.v += n
		tmpV.name += "add n .."
		imh.Change(tmpIndex, tmpV)
	}
	imh.Change(0, Item{math.MaxInt32, strconv.Itoa(math.MaxInt32)})
	if imh.Max().v != math.MaxInt32 {
		t.Fatal("IndexMapHeap change function not working")
	}
}
