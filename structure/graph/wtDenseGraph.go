package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type WtDenseGraph struct {
	v, e      int
	isDirectd bool
	g         [][]*edge // 这里放的是权值
}

func NewWtDenseGraph(v, e int, isDirected bool) WtDenseGraph {
	tmp := WtDenseGraph{
		v:         v,
		e:         e,
		isDirectd: isDirected,
		g:         nil,
	}
	tmp.init()
	return tmp
}

func NewWtDenseGraphFromFile(path string) (WtDenseGraph, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return WtDenseGraph{}, err
	}

	nVer, nEdge := 0, 0
	dg := WtDenseGraph{}
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
			dg = NewWtDenseGraph(nVer, nEdge, false)
			continue
		}
		dg.addEdge(l, r, wt)
	}

	if dg.e != totalEdge {
		return WtDenseGraph{}, fmt.Errorf("edge not equal")
	}
	err = scanner.Err()
	if err != nil {
		return WtDenseGraph{}, nil
	}

	return dg, nil
}

func (d *WtDenseGraph) init() {
	d.g = make([][]*edge, d.v)
	for i := 0; i < d.v; i++ {
		d.g[i] = make([]*edge, d.v)
	}
}

func (d *WtDenseGraph) E() int {
	return d.e
}

func (d *WtDenseGraph) V() int {
	return d.v
}

func (d *WtDenseGraph) addEdge(i, j int, wt float64) {
	if d.hasEdge(i, j) {
		d.g[i][j] = nil
		if !d.isDirectd {
			d.g[j][i] = nil
		}
		d.e--
	}

	d.g[i][j] = &edge{
		v:  i,
		w:  j,
		wt: wt,
	}
	if !d.isDirectd {
		d.g[j][i] = &edge{
			v:  j,
			w:  i,
			wt: wt,
		}
	}
	d.e++
}
func (d *WtDenseGraph) hasEdge(i, j int) bool {
	return d.g[i][j] != nil
}

func (d *WtDenseGraph) show() {
	for i := 0; i < d.V(); i++ {
		fmt.Print(i, ": ")
		tmpIter := NewWtDenseGraphIterator(d, i)
		for v := tmpIter.begin(); !tmpIter.end(); v = tmpIter.next() {
			if v != nil {
				fmt.Printf("%3.2f   ", v.wt)
			} else {
				fmt.Print("NULL   ")
			}
		}
		fmt.Println("")
	}
}
