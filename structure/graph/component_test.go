package graph

import (
	"fmt"
	"testing"
)

func TestCalComponentG1(t *testing.T) {
	sg1,err := NewSparseGraphFromFile("../../testG1.txt")
	if err != nil {
		t.Fatal("read graph from file error ", err.Error())
		return
	}

	com := NewComponent(&sg1, sg1.v)  // 注意这里为什么要传入指针， 因为只有指针实现了接口方法
	com.GetComponent()
	fmt.Printf("G1 component size :%d\n", com.count)
}

func TestCalComponentG2(t *testing.T) {
	sg2,err := NewDenseGraphFromFile("../../testG2.txt")
	if err != nil {
		t.Fatal("read graph from file error ", err.Error())
		return
	}

	com := NewComponent(&sg2, sg2.v)
	com.GetComponent()
	fmt.Printf("G2 component size :%d\n", com.count)
}


func TestCalComponentG3(t *testing.T) {
	sg2,err := NewDenseGraphFromFile("../../testG3")
	if err != nil {
		t.Fatal("read graph from file error ", err.Error())
		return
	}

	com := NewComponent(&sg2, sg2.v)
	com.GetComponent()
	fmt.Printf("G3 component size :%d\n", com.count)
}

func TestComponent_IsConnected(t *testing.T) {
	sg1,err := NewSparseGraphFromFile("../../testG1.txt")
	if err != nil {
		t.Fatal("read graph from file error ", err.Error())
		return
	}

	com := NewComponent(&sg1, sg1.v)
	com.GetComponent()
	if com.IsConnected(0, 1) != true {
		t.Log("is connected not working")
	}

	if com.IsConnected(0, 7) == true {
		t.Log("is connected not working")
	}
}