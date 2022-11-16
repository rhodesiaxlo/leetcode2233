package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

/**
循环
1. 进入初始循环的交件
2. 循环动力
3. 终止循环
*/
type Node struct {
	Val  int
	Next *Node
}

func main() {
	//fmt.Println("prime number ", countPrime(3))
	//eratosthenesSieve(3)
	fmt.Println(SieveOfEratosthenes(2))

	//deleteWhileIter()
	//deleteWhileIter2()
	//deleteWhileIter3()
	//deleteWhileIter4()
	//test()
	//forRange()

	//second := &Node{1, nil}
	//first := &Node{0,second}
	//third := first
	//second.Val = 10
	//
	//fmt.Println(third.Next.Val)

	//sortInReverseOrder()
	//fmt.Print(string(97) + "#")
	//typeConvert()
	//arrayOfDigToString()
	//sliceToArray()
	//fib := fibonacciGenerator()
	//fmt.Println("first fibo:",fib())
	//fmt.Println("second fibo:",fib())
	//fmt.Println("third fibo:",fib())
	//
	//i:=10
	//var i1 interface{}
	//i1 = i
	//switch j := i1.(type) {
	//case int :
	//	fmt.Println("type int, val :", j)
	//default:
	//	fmt.Println("no iead what type it is")
	//
	//}
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

func arrayOfDigToString() {
	iarr := []int{1, 2, 3, 4}
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(iarr), " ", "", -1), "[]"))

	farr := []float32{1.0, 2.0, 3.0, 4.0}
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(farr), " ", "", -1), "[]"))

	// write reduce function
}

func sliceToArray() {
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

func deleteWhileIter() {
	src := []int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range src {
		fmt.Print(v, " ")
		if v == 4 {
			src = append(src[:i], src[i+1:]...)
		}
	}
	fmt.Println()
	fmt.Println("cap:", len(src))
}
func deleteWhileIter4() {
	src := [...]int{1, 2, 3, 4, 5, 6, 7}
	for k, v := range src {
		fmt.Printf("k, v pointer of src is :%p\n", &src)
		fmt.Print(k, v)
	}

	fmt.Printf("origin pointer :%p", &src)
	fmt.Println()
	fmt.Println("cap:", len(src))
}

func deleteWhileIter2() {
	src := map[string]string{"name": "sennalu", "age": "20", "salary": "100"}
	for k, v := range src {
		fmt.Print(" k=", k, " v=", v)
		if k == "age" {
			delete(src, k)
		}
	}
	fmt.Println()
	fmt.Println("cap:", len(src))
}

/**
删除元素的正确做法
*/
func deleteWhileIter3() {
	src := []int{1, 2, 3, 4, 4, 6, 7}
	i := 0
	for {
		if i >= len(src) {
			break
		}

		if src[i] != 4 {
			i++
			continue
		}

		src = append(src[:i], src[i+1:]...)

	}
	fmt.Println(src)
	fmt.Println("cap:", len(src))
}

func test() {
	src := []int{1, 2, 3, 4}

	for i, j := 0, len(src); i < 4; i++ {
		if i < len(src) && src[i] == 2 {
			src = append(src[:i], src[i+1:]...)
		}
		fmt.Println("j=", j)
	}
}

/**
考点 for range 循环, 值专递， 本身，还是拷贝
*/
func forRange() {
	src := [...]int{1, 2, 3, 4, 5, 6}
	src2 := [...]int{1, 2, 3, 4, 5, 6}
	src3 := []int{1, 2, 3, 4, 5, 6}
	tmp := []int{}
	tmp2 := []int{}
	tmp3 := []int{}
	for k, v := range src {
		if k == 0 {
			src[1] = 22
			src[2] = 33
		}
		tmp = append(tmp, v)
	}
	fmt.Println("fix-length array for range example ")
	fmt.Println("original array ")
	fmt.Println(src)
	fmt.Println(tmp)

	for k, v := range src2[:] {
		if k == 0 {
			src2[1] = 22
			src2[2] = 33
		}
		tmp2 = append(tmp2, v)
	}
	fmt.Println("slice for range example, map is the same")
	fmt.Println("original array ")
	fmt.Println(src2)
	fmt.Println(tmp2)

	for k, v := range src3 {
		if k == 0 {
			src3[1] = 22
			src3[2] = 33
		}
		tmp3 = append(tmp3, v)
	}
	fmt.Println("slice for range example, map is the same")
	fmt.Println("original array ")
	fmt.Println(src3)
	fmt.Println(tmp3)
}

// 闭包可以包含逻辑 -- 可以和 interface 结合上
// middleware  http.HandleFunc("/hello", hello)
// func hello(w http.ResponseWriter, r *http.Request) {

func isPrime(x int) bool {
	if x == 1 {
		return false
	}

	if x == 2 {
		return true
	}

	i := 2
	for i < x {
		if x%i == 0 {
			return false
		}
		i++
	}

	return true
}

/**
计算 x 的范围内有多少个素数
*/
func countPrime(x int) int {
	ans := 0
	for i := 2; i <= x; i++ {
		if isPrime(i) {
			ans++
		}
	}

	return ans
}

func eratosthenesSieve(x int) {
	// do not use
	// prime := [num+1]bool
	// will cause : non-constant array bound num + 1 error

	// an array of boolean - the idiomatic way
	prime := make([]bool, x+1)

	// initialize everything with false first(not crossed)
	for i := 0; i < x+1; i++ {
		prime[i] = false
	}

	for i := 2; i*i <= x; i++ {
		if prime[i] == false {
			for j := i * 2; j <= x; j += i {
				prime[j] = true // cross
			}
		}

	}

}

func SieveOfEratosthenes(n int) []int {
	// Create a boolean array "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p*p <= n; p++ {
		// If integers[p] is not changed, then it is a prime
		if integers[p] == true {
			// Update all multiples of p
			for i := p * 2; i <= n; i += p {
				integers[i] = false
			}
		}
	}

	// return all prime numbers <= n
	var primes []int
	for p := 2; p <= n; p++ {
		if integers[p] == true {
			primes = append(primes, p)
		}
	}

	return primes
}

func evalOrder()  {
	i
}
