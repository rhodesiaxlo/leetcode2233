package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	Val int
	Next *Node
}
func main() {

	second := &Node{1, nil}
	first := &Node{0,second}
	third := first
	second.Val = 10

	fmt.Println(third.Next.Val)

	//sortInReverseOrder()
	//fmt.Print(string(97) + "#")
	//typeConvert()
	//arrayOfDigToString()
	sliceToArray()
	fib := fibonacciGenerator()
	fmt.Println("first fibo:",fib())
	fmt.Println("second fibo:",fib())
	fmt.Println("third fibo:",fib())
}

// 倒序排序
func sortInReverseOrder() {

	strs := []string{"flow", "fa", "flxw"}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] > strs[j]
	})
	fmt.Println(strs)

	ints := []int{1, 2, 3, 4, 5}
	sort.Slice(ints, func(i, j int) bool {
		return i > j
	})
	fmt.Println(ints)
}

func whileAndUntilInGo() {

}

func typeConvert() {
	// this is why u can not use  string(int) to convert int to string
	fmt.Println((string(97))) // a
	fmt.Printf("correct way to convert digits 123 to string is strconv.itoa %s \n", strconv.Itoa(123))

	// type basic operation
	i, err := strconv.ParseInt("123", 10, 32)
	if err != nil {
		fmt.Println("error , 123 can not be convert to on 2 base")
	} else {
		fmt.Println("123 parsed int is ", i)
	}

	i1 := 10
	fmt.Printf("defautl type of int is %v, len=%d", reflect.TypeOf(i1), bits.Len(uint(i)))
	// 1. check is type
	// 2. get type
	// 3. type conversion
	// 4. 调节


	// convert between int8 int16 int int32 int64
	// convert between float32 float64

}

/*
切片拓展的一个坑是如果 append(src, des) ， src 有容量会修改原有的数据
所以在
*/
func sliceAppend() {

}


func arrayOfDigToString()  {
	iarr := []int {1,2,3,4,}
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(iarr),  " ", "", -1), "[]"))

	farr := []float32 {1.0,2.0,3.0,4.0,}
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(farr),  " ", "", -1), "[]"))

	// write reduce function
}

func sliceToArray()  {
	var x []int
	for i := 0; i < 100; i++ {
		x = append(x, i)
	}

	var y [10]int
	copy(y[:], x[:11])
	fmt.Println(y)

	// after go 1.6， 可以直接指针转换
}

// 利用闭包实现feibonachi 数列
func fibonacciGenerator() func() int {
	a, b := 1, 1
	return func() int {
		a, b = a+b, a
		return a
	}
}

// 闭包可以包含逻辑 -- 可以和 interface 结合上
// middleware  http.HandleFunc("/hello", hello)
// func hello(w http.ResponseWriter, r *http.Request) {



