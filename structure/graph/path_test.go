package graph

import (
	"fmt"
	"testing"
)

func TestPath(t *testing.T) {
	sg1, err := NewSparseGraphFromFile("../../testG3")
	if err != nil {
		t.Fatal("read graph from file error ", err.Error())
		return
	}

	sg1.show()
	path := NewPath(&sg1, 0, sg1.v) // 注意这里为什么要传入指针， 因为只有指针实现了接口方法
	path.CalcPath()
	fmt.Print("0到5的路径\n")
	path.ShowPath(5)
	fmt.Print("0到5的路径\n")
	path.ShowPath(6)
}
