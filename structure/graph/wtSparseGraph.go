package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 添加权重信息, 为最短路径等更复杂的算法开路

type WtSparseGraph struct {
	v, e      int
	isDirectd bool
	g         [][]*edge // 考虑到 DenseGraph 可能会有没有变的情况，这里需要传指针
}

func NewWtSparseGraph(v, e int, isDirected bool) WtSparseGraph {
	tmp := WtSparseGraph{
		v:         v,
		e:         e,
		isDirectd: isDirected,
		g:         nil,
	}
	tmp.init()
	return tmp
}

func (d *WtSparseGraph) init() {
	d.g = make([][]*edge, d.v)
}

func NewWtSparseGraphFromFile(path string) (WtSparseGraph, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return WtSparseGraph{}, err
	}

	nVer, nEdge := 0, 0
	sg := WtSparseGraph{}
	scanner := bufio.NewScanner(file)
	totalEdge :=0
	for scanner.Scan() {
		tmpFields := strings.Fields(scanner.Text())
		l, _ := strconv.Atoi(tmpFields[0])
		r, _ := strconv.Atoi(tmpFields[1])
		wt := 0.0
		if len(tmpFields) > 2{
			wt, _ = strconv.ParseFloat(tmpFields[2], 64)
		}
		if nVer == 0 && nEdge == 0 {
			nVer = l
			//nEdge = r
			totalEdge = r
			sg = NewWtSparseGraph(nVer, nEdge, false)
			continue
		}
		sg.addEdge(l, r, wt)
	}

	if totalEdge != sg.e {
		return WtSparseGraph{}, fmt.Errorf("edge number not equal")
	}
	err = scanner.Err()
	if err != nil {
		return WtSparseGraph{}, nil
	}

	return sg, nil
}

func (d *WtSparseGraph) E() int {
	return d.e
}

func (d *WtSparseGraph) V() int {
	return d.v
}

func (d *WtSparseGraph) addEdge(i, j int, wt float64) {
	// 有可能是重复的，只是权重不同，这个时候需要删除原来的边，更新上新的边
	if d.hasEdge(i, j) {
		removeBoth := false
		if !d.isDirectd {
			removeBoth = true
		}
		d.removeEdge(i, j, removeBoth)
	}

	d.g[i] = append(d.g[i], &edge{
		v:  i,
		w:  j,
		wt: wt,
	})
	if i != j && !d.isDirectd { // 1. 同时排除了自环边 self-loop,
		d.g[j] = append(d.g[j], &edge{
			v:  j,
			w:  i,
			wt: wt,
		})
	}
	d.e++
}
func (d *WtSparseGraph) hasEdge(i, j int) bool {
	for k := 0; k < len(d.g[i]); k++ {
		tmpOther, _ := d.g[i][k].Other(i)
		if tmpOther == j {
			return true
		}
	}
	return false
}

func (d *WtSparseGraph) removeEdge(i, j int, both bool)  {
	for k := 0; k < len(d.g[i]); k++ {
		tmpOther, _ := d.g[i][k].Other(i)
		if tmpOther == j {
			d.g[i] = append(d.g[i][:k], d.g[i][k+1:]...)
		}
	}

	if both {
		for l := 0; l < len(d.g[j]); l++ {
			tmpOther2, _ := d.g[j][l].Other(j)
			if tmpOther2 == i {
				  d.g[j] = append(d.g[j][:l], d.g[j][l+1:]...)
				//d.g[j] = append(d.g[j][:l], d.g[i][l+1:]...)
			}
		}
	}

	// 边数减1
	d.e--
}

func (d *WtSparseGraph) show() {
	for i := 0; i < d.V(); i++ {
		fmt.Print(i, ": ")
		tmpIter := NewWtIterator(d, i)
		for v := tmpIter.begin(); !tmpIter.end(); v = tmpIter.next() {
			otherJ, _ := v.Other(i)
			fmt.Printf("(to:%-4d wt:%3.2f)  ", otherJ, v.wt)
		}
		fmt.Println("")
	}
}