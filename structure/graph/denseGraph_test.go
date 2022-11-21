package graph

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestCreateDenseGraph(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	dg := NewDenseGraph(N, 0, false)
	for i := 0; i < M; i++ {
		j := r.Intn(N)
		k := r.Intn(N)
		dg.addEdge(j, k)
	}

	// 遍历
	fmt.Println("dense graph traverse ------------")
	for i := 0; i < N; i++ {
		fmt.Print(i, ": ")
		for j := 0; j < N; j++ {
			if dg.g[i][j] == true {
				fmt.Print(j, " ")
			}
		}
		fmt.Println("")
	}
}

func TestReadDenseGraphFromFile(t *testing.T) {
	file, err := os.Open("../../testG1.txt")
	defer file.Close()
	if err != nil {
		t.Log("open file error ", err.Error())
		return
	}

	edge := [][2]int(nil)
	nVer, nEdge := 0, 0
	sg := DenseGraph{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmpFields := strings.Fields(scanner.Text())
		l, _ := strconv.Atoi(tmpFields[0])
		r, _ := strconv.Atoi(tmpFields[1])
		if nVer == 0 || nEdge == 0 {
			nVer = l
			nEdge = r
			sg = NewDenseGraph(nVer, nEdge, false)
			continue
		}
		sg.addEdge(l, r)
	}
	err = scanner.Err()
	if err != nil {
		t.Fatal("read data from filer err ", err.Error())
		return
	}

	for _, v := range edge {
		sg.addEdge(v[0], v[1])
	}

	// traverse
	for i := 0; i < sg.v; i++ {
		fmt.Print(i, ": ")
		for j := 0; j < len(sg.g[i]); j++ {
			if sg.g[i][j] == true {
				fmt.Print(j, " ")
			}
		}
		fmt.Println("")
	}
}

func TestIterDenseUsingIterator(t *testing.T) {
	dg, _ := NewDenseGraphFromFile("../../testG1.txt")
	for i := 0; i < dg.V(); i++ {
		fmt.Print(i, ": ")
		tmpIter := NewDenseGraphIterator(&dg, i)
		for v := tmpIter.begin(); !tmpIter.end(); v = tmpIter.next() {
			fmt.Print(v, " ")
		}
		fmt.Println("")
	}
}

func TestIter2(t *testing.T) {
	dg, _ := NewDenseGraphFromFile("../../testG1.txt")
	dg.show()
}
