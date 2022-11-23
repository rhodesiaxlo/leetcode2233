package graph

import "fmt"

// 边，增加权重信息

type edge struct {
	v, w int
	wt   float64
}

func NewEdge(v, w int, wt float64)  *edge {
	return &edge{
		v:  v,
		w:  w,
		wt: wt,
	}
}

func (e *edge) V() int {
	return e.v
}

func (e *edge) W() int {
	return e.w
}

func (e *edge) Weight() float64 {
	return e.wt
}

func (e *edge) Other(x int) (int, error) {
	if x != e.v && x != e.w {
		return 0, fmt.Errorf("index illegal")
	}

	if x == e.v {
		return e.w, nil
	}

	return e.v,nil
}

