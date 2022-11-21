package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
流程

1. 建立模型
2. 建立连接
3. 遍历
4. 其它算法
*/
type DenseGraph struct {
	v, e      int
	isDirectd bool
	g         [][]bool
}

func NewDenseGraph(v, e int, isDirected bool) DenseGraph {
	tmp := DenseGraph{
		v:         v,
		e:         e,
		isDirectd: isDirected,
		g:         nil,
	}
	tmp.init()
	return tmp
}

func NewDenseGraphFromFile(path string) (DenseGraph,error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return DenseGraph{}, err
	}

	edge := [][2]int(nil)
	nVer, nEdge := 0, 0
	dg := DenseGraph{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmpFields := strings.Fields(scanner.Text())
		l, _ := strconv.Atoi(tmpFields[0])
		r, _ := strconv.Atoi(tmpFields[1])
		if nVer == 0 || nEdge == 0 {
			nVer = l
			nEdge = r
			dg = NewDenseGraph(nVer, nEdge, false)
			continue
		}
		dg.addEdge(l, r)
	}
	err = scanner.Err()
	if err != nil {
		return DenseGraph{}, nil
	}

	for _, v := range edge {
		dg.addEdge(v[0], v[1])
	}
	return dg, nil
}

func (d *DenseGraph) init() {
	d.g = make([][]bool, d.v)
	for i := 0; i < d.v; i++ {
		d.g[i] = make([]bool, d.v)
	}
}

func (d *DenseGraph) E() int {
	return d.e
}

func (d *DenseGraph) V() int {
	return d.v
}

func (d *DenseGraph) addEdge(i, j int) {
	if d.hasEdge(i, j) {
		return
	}

	d.g[i][j] = true // 没有平行边
	if !d.isDirectd {
		d.g[j][i] = true
	}
	d.e++
}
func (d *DenseGraph) hasEdge(i, j int) bool {
	if d.isDirectd {
		return d.g[i][j]
	} else {
		return d.g[i][j] && d.g[j][i]
	}
}

func (d *DenseGraph) show()  {
	for i:=0;i<d.V();i++ {
		fmt.Print(i, ": ")
		tmpIter := NewDenseGraphIterator(d, i)
		for v:=tmpIter.begin();!tmpIter.end();v=tmpIter.next() {
			fmt.Print(v, " ")
		}
		fmt.Println("")
	}
}

