package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type SomeItem struct {
	i int
	j []int
}

func main() {
	var s1 []int // nil 切片
	fmt.Println(s1, *(*reflect.SliceHeader)(unsafe.Pointer(&s1)))
	s2 := []int(nil) // nil 切片
	fmt.Println(s2, *(*reflect.SliceHeader)(unsafe.Pointer(&s2)))
	s3 := make([]int, 0) // 空切片
	fmt.Println(s3, *(*reflect.SliceHeader)(unsafe.Pointer(&s3)))
	s4 := make([]int, 2, 5) // 零切片
	fmt.Println(s4, *(*reflect.SliceHeader)(unsafe.Pointer(&s4)))
	s5 := []int{} //
	s6 := []int(nil)
	fmt.Print(s5, s6)

	s7 := make([]int, 4, 8)
	for i := 0; i < 4; i++ {
		s7[i] = i
	}
	fmt.Println("------------------------------------------")
	fmt.Println(s7, *(*reflect.SliceHeader)(unsafe.Pointer(&s7)))
	// out of length 就会outof bound
	// fmt.Printf("s7[0]= %d, s7[1]=%d, s7[2]=%d, s7[3]=%d, s7[4]=%d \n", s7[0], s7[1], s7[2], s7[3], s7[4])  	s7[4] 会 out of bound
	s8 := s7[1:2]
	fmt.Println(s8, *(*reflect.SliceHeader)(unsafe.Pointer(&s8)))
	// fmt.Printf("s8[0]= %d, s8[1]=%d \n", s8[0], s8[1]) // s8[1] 会 outof bound
	// slice 的 slice 会重新分配 length

	// slice 的扩容
	//s9 := []int(nil)
	//preCap := 0
	//for i := 0; i < 1000000; i++ {
	//	s9 = append(s9, i)
	//	tmpSlice := *(*reflect.SliceHeader)(unsafe.Pointer(&s9))
	//	if tmpSlice.Cap > preCap {
	//		fmt.Printf("len:%d, cap:%d\n", tmpSlice.Len, tmpSlice.Cap)
	//		preCap = tmpSlice.Cap
	//	}
	//}

	// 头部删除切片和尾部删除切片
	s10 := make([]int, 100)
	fmt.Println("remove head of tail from slice", *(*reflect.SliceHeader)(unsafe.Pointer(&s10)))
	// 切除尾巴, 还是原来的地址
	s11 := s10[:2]
	fmt.Println("remove tail from slice", *(*reflect.SliceHeader)(unsafe.Pointer(&s11)))
	// 切除头部， 已经重新开辟地址
	s12 := s10[2:3]
	fmt.Println("remove head from slice", *(*reflect.SliceHeader)(unsafe.Pointer(&s12)))

	// 拷贝切片， 深度拷贝
	s13 := []SomeItem(nil)
	s13 = append(s13, SomeItem{
		i: 1,
		j: []int{1, 2, 3, 4},
	})
	s13 = append(s13, SomeItem{
		i: 2,
		j: []int{5, 6, 7, 8},
	})

	s14 := s13[1:2]
	s15 := make([]SomeItem, 1)
	copy(s15, s13[1:2])
	(s13[1]).j = []int(nil)
	(s13[1]).i = 100
	fmt.Println("s13, origin slice, wait to be sliced or copy", s13[0])
	fmt.Println("s14, assign", s14[0]) // assign 虽然是拷贝，但是内容没有变
	fmt.Println("s15, copy ", s15[0])  // 深度拷贝，i 没有收到影响
	// 如果 slice 的内容是 *SomeItem 而不是 SomeItem ,就无回天之力

	// copy nil slice
	// copy 4 number slice to 2 number slice
	// copy 4 nmber slice to 10 number slice
	s16 := make([]int, 8)
	for i := 0; i < len(s16); i++ {
		s16[i] = i
	}
	s17 := []int(nil)
	copy(s17, s16)
	s18 := make([]int, 4)
	copy(s18, s16)
	s19 := make([]int, 10)
	copy(s19, s16)
	fmt.Println("s16 to be copyied, length of 8 , we test to copy to nil slice, slice of length 4, slice of length 10")
	fmt.Println("s16:", s16)
	fmt.Println("s17:", s17)
	fmt.Println("s18:", s18)
	fmt.Println("s19:", s19)

	// slice 相等
	//sliceEqual()

	fmt.Println("---------")
	s21 := [...]int{0: 1, 2: 1, 3: 4}
	fmt.Println("s20", *(*reflect.SliceHeader)(unsafe.Pointer(&s21)), s21[3], "len=", len(s21))

	// append 后是否有必要 assign 呢个
	s22 := make([]int, 3)
	for i22:=0;i22< 3;i22++ {
		s22[i22] = i22
	}
	fmt.Printf("s22 to extend , origin address is %p\n", s22)
	s23 := make([]int, 1000)
	for i23:=0;i23< 1000;i23++ {
		s23[i23] = 1000+i23
	}
	s24 := append(s22, s23...)
	fmt.Printf("s22 after extend , origin address is %p  v:%v\n", s22, s22)  // 所以 append 之后一定要返回extend 值
	fmt.Printf("s24 ptr:=%p  v:%v\n", s24, s24[:10])
}

func sliceEqual(s1, s2 []int) {
	// 创建两个长度相同、容量不同的切片
	//slice1 := make([]string,2,5)
	//slice1[0] = "Go"
	//slice1[1] = "语"
	//slice2 := make([]string,2,6)
	//slice2[0] = "Go"
	//slice2[1] = "语"
	//// 方法1：自定义方法
	//status := true
	//if len(slice1) != len(slice2){
	//	status = false
	//}else{
	//	for k,_ := range slice1 {
	//		if slice1[k] != slice2[k] {
	//			status = false
	//		}
	//	}
	//}
	//fmt.Println(status)     // 输出：true
	//// 方法2：使用 reflect.DeepEqual
	//fmt.Println(reflect.DeepEqual(slice1,slice2))  // 输出：true
}
