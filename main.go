package main

import (
	"bufio"
	"fmt"
	"leetcode/leetcode/algo"
	"leetcode/leetcode/sorting"
	"leetcode/leetcode/sorting/linearSort"
	"leetcode/leetcode/structure"
	"leetcode/leetcode/structure/heap"
	"leetcode/leetcode/structure/list"
	"leetcode/leetcode/structure/tree"
	"math/bits"
	"math/rand"
	"os"
	"reflect"
	"regexp"
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
	printNonCompleteTree()
	return
	//fmt.Println(nil == nil)
	//return

	x100 := []int(nil)
	//x1 := []int{}
	x100  = append(x100, 1)
	fmt.Println(x100)
	return

	bst3 := tree.NewBst()
	bst3.Insert("20", 20)
	bst3.Insert("10", 10)
	bst3.Insert("30", 30)
	bst3.Insert("5", 5)
	bst3.Insert("15", 15)
	bst3.Insert("100", 100)
	bst3.LevelOrder()
	return
	// 20 10 30 15 5 100
	bst2 := tree.NewBst()
	for i:=0;i < 10;i++ {
		tmpV := rand.Intn(10000)
		bst2.Insert(strconv.Itoa(tmpV), tmpV)
	}
	bst2.InOrder()
	return

	bst := tree.NewBst()
	file, err := os.Open("./bible3.txt")
	if err != nil {
		fmt.Println("opps , we have a problem", err.Error())
		return
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)
	total := 0
	for scan.Scan() {
		total++
		tmpWord := scan.Text()
		tmpWord1 := strings.ToLower(tmpWord)
		compile, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
		tmpWord1 = string(compile.ReplaceAll([]byte(tmpWord1), []byte("")))
		e := bst.Search(tmpWord1)
		if e != nil {
			e.V++
		} else {
			bst.Insert(tmpWord1, 1)
		}
	}
	file.Close()
	god := bst.Search("god")


	fmt.Print("god number", god.V, "total ", total, " bst total ", bst.Agg())
	return

	nums1 := []int{1, 2, 3, 4, 5}
	changeVal(nums1)
	nums3 := []int{1, 2, 3, 4, 5} // !!! replace
	changeValByAssign(nums3)

	fmt.Println("nums1", nums1)
	fmt.Println("nums3", nums3)
	return

	arr := []int{1, 2, 3, 4, 5, 6}
	arr2 := arr[:3]
	fmt.Println(len(arr2))
	return
	nums := sorting.GenArray(10, 0)
	mh := heap.MaxHeap{}
	mh.Heapify(nums)
	linearSort.QuickSort(nums)
	max := nums[len(nums)-1]
	fmt.Println(max)
	return

	heapArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	structure.PrintComTree(heapArr)
	//fmt.Println(arr)
	return

	str := make([]string, 3)
	fmt.Println(str)
	return

	a := [][2]int{{1, 2}, {3, 4}, {5, 6}}

	fmt.Println(a)
	return
	nums2 := sorting.GenArrayDesc(10)
	linearSort.InsertSort(nums2)
	fmt.Println(nums2)
	return
	//nums := sorting.GenArrayDesc(10)
	linearSort.MergeSort(nums)
	return
	s1 := rand.NewSource(10)
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Float64())
	fmt.Println(r1.Float64())
	fmt.Println(r1.Float64())
	fmt.Println("random:", r1.Int())
	fmt.Println("random:", r1.Int())
	fmt.Println("-------------------  same source, same random number")
	s2 := rand.NewSource(10)
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Float64())
	fmt.Println(r2.Float64())
	fmt.Println(r2.Float64())
	fmt.Println("random:", r1.Int())
	fmt.Println("random:", r1.Int())

	return

	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	return
	//s := "abca"
	//fmt.Print(s[1:2]))
	x := []string{"a", "b", "c"}
	fmt.Println(reflect.TypeOf(x[:1]))
	return
	i := 0
	fmt.Scanf("%d", &i)
	fmt.Print("you have input ", i)
	//r := bufio.NewReader(os.Stdin)
	//s, _, err := r.ReadLine()
	//if err != nil {
	//	fmt.Println("input error")
	//}
	//
	//cnt := 0
	//for i:= len(s)-1;i >= 0 ;i-- {
	//	if s[i] == ' ' {
	//		break
	//	}
	//	cnt++
	//}
	//fmt.Print(cnt)
}

func main1() {
	bst2 := tree.NewBst()
	for i:=0;i < 10;i++ {
		tmpV := rand.Intn(10000)
		bst2.Insert(strconv.Itoa(tmpV), tmpV)
	}
	bst2.InOrder()
	return

	algo.LeetCode32longestValidParentheses("))(()()))()()()()")
	algo.LeetCode28IndexOfStr("leetcode", "leeto")
	n1 := &list.ListNode{
		Val:  0,
		Next: nil,
	}
	n2 := &list.ListNode{
		Val:  1,
		Next: nil,
	}
	n1.Next = n2
	n3 := &list.ListNode{
		Val:  3,
		Next: nil,
	}
	n2.Next = n3

	//fmt.Println(algo.LeetCode876MiddleOfLinkList(&list.ListNode{}))

	for i := 3; forCond(i); i-- {

	}

	fmt.Println(algo.LeetCode29divide(10, 3))

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

func leetcode876() {
	n1 := &list.ListNode{
		Val:  0,
		Next: nil,
	}
	n2 := &list.ListNode{
		Val:  1,
		Next: nil,
	}
	n1.Next = n2
	n3 := &list.ListNode{
		Val:  3,
		Next: nil,
	}
	n2.Next = n3

	fmt.Println(algo.LeetCode876MiddleOfLinkList(&list.ListNode{}))
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

func evalOrder() {

}

func forCond(n int) bool {
	fmt.Println(n)
	return n > 0
}

func changeVal(nums []int) {
	nums[0] = 1000
}

func changeValByAssign(nums []int) {
	nums1 := make([]int, len(nums))
	copy(nums1, nums)
	nums = nums1 // replace
}

func printNonCompleteTree() {
	root := &tree.TreeNode{
		K: "0",
		V: 0,
		L: nil,
		R: nil,
	}

	level1l1 := &tree.TreeNode{
		K: "1",
		V: 1,
		L: nil,
		R: nil,
	}

	level1r1 := &tree.TreeNode{
		K: "2",
		V: 2,
		L: nil,
		R: nil,
	}

	root.L = level1l1
	root.R = level1r1

	level2L1 := &tree.TreeNode{
		K: "3",
		V: 3,
		L: nil,
		R: nil,
	}
	level2r1 := &tree.TreeNode{
		K: "4",
		V: 4,
		L: nil,
		R: nil,
	}
	level2r2 := &tree.TreeNode{
		K: "5",
		V: 5,
		L: nil,
		R: nil,
	}
	level1l1.L = level2L1
	level1r1.L = level2r1
	level1r1.R = level2r2
	level3r1 := &tree.TreeNode{
		K: "6",
		V: 6,
		L: nil,
		R: nil,
	}
	level2r2.L = level3r1

	structure.PrintNonCompleteTree(root)
}
