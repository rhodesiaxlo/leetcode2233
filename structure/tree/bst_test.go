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
	n := 1000
	m := make(map[int]bool)
	bst := NewBst()
	for i := 0; i < n; i++ {
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

func TestGetMinAndMax(t *testing.T) {
	n := 30
	m := make(map[int]bool)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	bst := NewBst()
	min := "ZZZZZZZZZZZZZZZZ"
	max := ""
	for i := 0; i < n; i++ {
		tmpInt := r.Intn(100000)
		if v, ok := m[tmpInt]; !ok || !v {
			tmpStr := strconv.Itoa(tmpInt)
			if tmpStr > max {
				max = tmpStr
			}

			if tmpStr < min {
				min = tmpStr
			}
			bst.Insert(strconv.Itoa(tmpInt), tmpInt)
			m[tmpInt] = true
		}
	}

	minNode := bst.Minimum()
	maxNode := bst.Maximum()
	if minNode.K != min {
		t.Fatal("bst minimum function not working")
	}

	if maxNode.K != max {
		t.Fatal("bst maximum function not working")
	}
}

func TestRemoveMin(t *testing.T) {
	n := 10
	m := make(map[int]bool)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	bst := NewBst()
	min := "ZZZZZZZZZZZZZZZZ"
	max := ""
	for i := 0; i < n; i++ {
		tmpInt := r.Intn(100000)
		if v, ok := m[tmpInt]; !ok || !v {
			tmpStr := strconv.Itoa(tmpInt)
			if tmpStr > max {
				max = tmpStr
			}

			if tmpStr < min {
				min = tmpStr
			}
			bst.Insert(strconv.Itoa(tmpInt), tmpInt)
			m[tmpInt] = true
		}
	}

	var minStr []string
	for bst.Minimum() != nil {
		minStr = append(minStr, bst.RemoveMin().K)
	}

	for i := 0; i < len(minStr)-1; i++ {
		if minStr[i] > minStr[i+1] {
			t.Fatal("remove min function not working")
		}
	}
}

func TestRemoveMax(t *testing.T) {
	n := 10
	m := make(map[int]bool)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	bst := NewBst()
	min := "ZZZZZZZZZZZZZZZZ"
	max := ""
	for i := 0; i < n; i++ {
		tmpInt := r.Intn(100000)
		if v, ok := m[tmpInt]; !ok || !v {
			tmpStr := strconv.Itoa(tmpInt)
			if tmpStr > max {
				max = tmpStr
			}

			if tmpStr < min {
				min = tmpStr
			}
			bst.Insert(strconv.Itoa(tmpInt), tmpInt)
			m[tmpInt] = true
		}
	}

	var maxStr []string
	for bst.Maximum() != nil {
		maxStr = append(maxStr, bst.RemoveMax().K)
	}

	for i := 0; i < len(maxStr)-1; i++ {
		if maxStr[i] < maxStr[i+1] {
			t.Fatal("remove min function not working")
		}
	}
}

func TestHibardDeletion(t *testing.T) {
	xx := []string{
		"50", "30", "70", "15", "43", "10", "11", "40", "44", "55", "52", "56", "75", "71",
	}

	/*
					      50
					/            \
				   30            70
	             /   \          /   \
				 15   43       55    75
				/ \	  /	\      /\     /\
		       10    40 44   52 56  71
				\11
	*/
	bst := NewBst()
	for _, v := range xx {
		tmpInt, _ := strconv.Atoi(v)
		bst.Insert(v, tmpInt)
	}

	// count
	if bst.Size() != len(xx) {
		t.Fatal("size function not working")
	}

	bst.Remove("70")
	if bst.Size() !=len(xx)-1 {
		t.Fatal("size function not working")
	}

	bst.Remove("30")
	if bst.Size() !=len(xx)-2 {
		t.Fatal("size function not working")
	}
}

/**
bst 顺序插入在10k 的数据规模下性能很低
 */
func TestBstInsertByOrder(t *testing.T) {
	// todo
	// communist.txt
}

func TestTrieVsBst(t *testing.T) {
	// todo
	// bible.txt
}