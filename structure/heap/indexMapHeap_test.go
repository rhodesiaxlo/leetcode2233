package heap

import (
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// 测试insert
func TestIndexMaxHeap_Insert(t *testing.T) {
	n:=1000000
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}

	imh := IndexMaxHeap{}
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
			t.Fatal("IndexMapHeap insert sort not working")
		}
	}
}

// heapify
func TestIndexMaxHeap_Heapify(t *testing.T) {
	n:=1000000
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}
	imh := IndexMaxHeap{}
	imh.Heapify(items)
	imh.SortAsc()
	ans := imh.GetArr()
	sz2 := len(ans)
	for i:=0;i < sz2-1;i++ {
		if ans[i].Gt(ans[i+1]) {
			t.Fatal("IndexMapHeap heapify sort not working")
		}
	}
}

func TestIndexMaxHeap_Max(t *testing.T) {
	n:=1000000
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}
	imh := IndexMaxHeap{}
	imh.Heapify(items)
	if imh.Max().v != n-1 {
		t.Fatal("IndexMaxHeap max function not working")
	}
}

// insert
func TestIndexMaxHeap_InsertAndSort(t *testing.T) {
	n:=1000000
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}
	imh := IndexMaxHeap{}
	imh.Heapify(items)
	imh.Insert(Item{
		v:    n,
		name: strconv.Itoa(n),
	})
	if imh.Max().v != n {
		t.Fatal("IndexMapHeap insert after Heapify not working")
	}
}

// change
func TestIndexMaxHeap_Change(t *testing.T) {
	n:=1000000
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}
	imh := IndexMaxHeap{}
	imh.Heapify(items)
	imh.Change(0, Item{math.MaxInt32, strconv.Itoa(math.MaxInt32)})
	if imh.Max().v != math.MaxInt32 {
		t.Fatal("IndexMapHeap change function not working")
	}
}

// getItems
func TestIndexMaxHeap_GetItem(t *testing.T) {
	n:=1000000
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}
	imh := IndexMaxHeap{}
	tmpItem, _:= imh.GetItem(0)
	if tmpItem.v != 0 {
		t.Fatal("IndexMapHeap getItem not working")
	}

}

func TestIndexMaxHeap_ChangeSoManyTimes(t *testing.T) {
	n:=1000000
	var items []Item
	for i:=0;i<n;i++ {
		items = append(items, Item{
			v:    i,
			name: strconv.Itoa(i),
		})
	}

	imh := IndexMaxHeap{}
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

