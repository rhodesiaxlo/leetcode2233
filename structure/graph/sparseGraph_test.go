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

const N = 20
const M = 100

func TestCreateSparseGraph(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	sg := NewSparseGraph(N, 0, false)
	for i := 0; i < M; i++ {
		j := r.Intn(N)
		k := r.Intn(N)
		sg.addEdge(j, k)
	}

	// 遍历
	fmt.Println("sparse graph traverse ------------")
	for i := 0; i < N; i++ {
		fmt.Print(i, ": ")
		for j := 0; j < len(sg.g[i]); j++ {
			fmt.Print(sg.g[i][j], " ")
		}
		fmt.Println("")
	}
}


func TestReadSparseGraphFromFile(t *testing.T) {
	file, err := os.Open("../../testG1.txt")
	defer file.Close()
	if err != nil {
		t.Log("open file error ", err.Error())
		return
	}

	edge := [][2]int(nil)
	nVer, nEdge := 0, 0
	sg := SparseGraph{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmpFields := strings.Fields(scanner.Text())
		l, _ := strconv.Atoi(tmpFields[0])
		r, _ := strconv.Atoi(tmpFields[1])
		if nVer == 0 || nEdge == 0 {
			nVer = l
			nEdge = r
			sg = NewSparseGraph(nVer, nEdge, false)
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
	for i:=0;i<sg.v;i++ {
		fmt.Print(i,": ")
		for j:=0;j<len(sg.g[i]);j++ {
			fmt.Print(sg.g[i][j], " ")
		}
		fmt.Println("")
	}
}

func TestIterUsingIterator(t *testing.T) {
	sg, _ := NewSparseGraphFromFile("../../testG1.txt")
	fmt.Println("sparse graph traverse ------------")
	for i := 0; i < sg.V(); i++ {
		fmt.Print(i, ": ")
		tmpIterator := NewIterator(&sg, i)
		for v := tmpIterator.begin(); !tmpIterator.end(); v = tmpIterator.next() {
			fmt.Print(v, " ")
		}
		fmt.Println("")
	}
}
