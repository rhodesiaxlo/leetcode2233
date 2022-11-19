package tree

import (
	"bufio"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestBst_Insert(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n :=1000
	m := make(map[int]bool)
	bst := NewBst()
	for i:=0;i<n;i++ {
		tmp := r.Intn(math.MaxInt32)
		if v, ok := m[tmp]; !ok || !v {
			bst.Insert(strconv.Itoa(tmp), tmp)
			m[tmp] = true
		}
	}

	if bst.Size() != n {
		t.Fatal("bst size funciton not working")
	}
}

func TestBstInsert(t *testing.T) {
	file, err := os.Open("../../bible.txt")
	if err != nil {
		t.Fatal("file open error ", err.Error())
	}

	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)
	bst := NewBst()
	for scan.Scan() {
		w := strings.ToLower(scan.Text())
		exist := bst.Search(w)
		if exist != nil {
			exist.V++
		} else {
			bst.Insert(w, 1)
		}
	}
	file.Close()
}

/**
作为对比 40w 数据量下的性能差异
 */
func TestSst_Insert(t *testing.T) {
	file, err := os.Open("../../bible.txt")
	if err != nil {
		t.Fatal("file open error ", err.Error())
	}

	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)
	sst := NewSst()
	for scan.Scan() {
		w := strings.ToLower(scan.Text())
		exist := sst.Search(w)
		if exist != nil {
			exist.V++
		} else {
			sst.Insert(w, 1)
		}
	}
}




