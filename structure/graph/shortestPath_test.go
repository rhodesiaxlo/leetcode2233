package graph

import (
	"testing"
)

func TestNewShortestPath(t *testing.T) {
	sg1, err := NewSparseGraphFromFile("../../testG3")
	if err != nil {
		t.Fatal("read graph from file error ", err.Error())
		return
	}

	sg1.show()
	path := NewShortestPath(&sg1, 0, sg1.v) // 注意这里为什么要传入指针， 因为只有指针实现了接口方法
	path.ShowPath(6)
}
