package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
graph 的分类

1. cycly graph  acyclic graph
2. weighted graph / unweighted graph
3. directed graph/ undirected graph
4. sparse graph /dense graph

4. self-loop  parallel-edges
5. simple graph


组成
edge  vertex
component
graph


adjacency matric
adgacency list

*/

type SparseGraph struct {
	v, e      int
	isDirectd bool
	g         [][]int // 记录的不再是 boll, 而是 int
}

func NewSparseGraph(v, e int, isDirected bool) SparseGraph {
	tmp := SparseGraph{
		v:         v,
		e:         e,
		isDirectd: isDirected,
		g:         nil,
	}
	tmp.init()
	return tmp
}

func NewSparseGraphFromFile(path string) (SparseGraph, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return SparseGraph{}, err
	}

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
		return SparseGraph{}, nil
	}

	return sg, nil
}

func (d *SparseGraph) init() {
	d.g = make([][]int, d.v)
	//for i := 0; i < d.v; i++ {
	//	d.g[i] = make([]int, d.v)
	//}
}

func (d *SparseGraph) E() int {
	return d.e
}

func (d *SparseGraph) V() int {
	return d.v
}

func (d *SparseGraph) addEdge(i, j int) {
	//if d.hasEdge(i, j) {  // 2. 因为性能的考虑， 没有判断是否已经有边， 后面统一处理重复
	//	return
	//}

	d.g[i] = append(d.g[i], j)
	if i != j && !d.isDirectd { // 1. 同时排除了自环边 self-loop,
		d.g[j] = append(d.g[j], i)
	}
	d.e++
}
func (d *SparseGraph) hasEdge(i, j int) bool {
	// 对于 adjacency list 为了判断是否有这条边，只能遍历，性能很低
	for k := 0; k < len(d.g[i]); k++ {
		if d.g[i][k] == j {
			return true
		}
	}
	return false
}

func (d *SparseGraph) show() {
	for i := 0; i < d.V(); i++ {
		fmt.Print(i, ": ")
		tmpIter := NewSparseGraphIterator(d, i)
		for v := tmpIter.begin(); !tmpIter.end(); v = tmpIter.next() {
			fmt.Print(v, " ")
		}
		fmt.Println("")
	}
}
