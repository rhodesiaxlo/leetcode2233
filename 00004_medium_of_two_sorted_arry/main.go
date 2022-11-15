package main

import (
	"fmt"
	"math"
)

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).
*/

func main() {
	num1 := [...]int{1, 2, 3, 4, 5}
	num2 := [...]int{6, 7, 8}

	fmt.Println(findMediumOfTwoSortedArray10(num1[:], num2[:]))
}

// 方法1
// 归并排序
func findMediumOfTwoSortedArray(nums1, nums2 []int) float64 {
	var l1, l2, i, j, k int

	l1 = len(nums1) // type convert
	l2 = len(nums2)
	m := make([]int, l1+l2)
	i = 0
	j = 0
	k = 0

	for {

		if i > l1-1 && j > l2-1 {
			break
		}

		if i == l1 {
			m[k] = nums2[j]
			j++
			k++
			continue
		}

		if j == l2 {
			m[k] = nums1[i]
			i++
			k++
			continue
		}

		if nums1[i] >= nums2[j] {
			m[k] = nums2[j]
			j++
			k++
			continue
		}

		if nums1[i] < nums2[j] {
			m[k] = nums1[i]
			i++
			k++
			continue
		}
	}
	fmt.Println(m)

	if (l1+l2)%2 == 1 {
		return float64(m[(l1+l2-1)/2])
	}

	return (float64(m[(l1+l2-1)/2]) + float64(m[(l1+l2+1)/2])) / 2
}

// 使用了固定数组取代切片， 减少了重新分配内存的开销
func findMediumOfTwoSortedArray2(nums1, nums2 []int) float64 {
	var l1, l2, i, j int

	l1 = len(nums1) // type convert
	l2 = len(nums2)
	var m []int
	i = 0
	j = 0

	for {

		if i > l1-1 && j > l2-1 {
			break
		}

		if i == l1 {
			m = append(m, nums2[j])
			j++
			continue
		}

		if j == l2 {
			m = append(m, nums1[i])
			i++
			continue
		}

		if nums1[i] >= nums2[j] {
			m = append(m, nums1[j])
			j++
			continue
		}

		if nums1[i] < nums2[j] {
			m = append(m, nums1[i])
			i++
			continue
		}
	}

	if (l1+l2)%2 == 1 {
		return float64(m[(l1+l2-1)/2])
	}

	return (float64(m[(l1+l2-1)/2]) + float64(m[(l1+l2+1)/2])) / 2
}

// 归并排序只排一部分
func findMediumOfTwoSortedArray3(nums1, nums2 []int) float64 {
	var l1, l2, i, j int

	l1 = len(nums1) // type convert
	l2 = len(nums2)
	var m []int
	i = 0
	j = 0
	max := (l1+l2)/2 + 1
	for {

		if i > l1-1 && j > l2-1 {
			break
		}

		if len(m) >= max {
			break
		}

		if i == l1 {
			m = append(m, nums2[j])
			j++
			continue
		}

		if j == l2 {
			m = append(m, nums1[i])
			i++
			continue
		}

		if nums1[i] >= nums2[j] {
			m = append(m, nums2[j])
			j++
			continue
		}

		if nums1[i] < nums2[j] {
			m = append(m, nums1[i])
			i++
			continue
		}
	}
	if (l1+l2)%2 == 1 {
		return float64(m[(l1+l2-1)/2])
	}

	return (float64(m[(l1+l2-1)/2]) + float64(m[(l1+l2+1)/2])) / 2
}

/**
分治法
 */
func findMediumOfTwoSortedArray10(nums1, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	if l2 > l1 {
		return findMediumOfTwoSortedArray10(nums2, nums1)
	}

	low := 0
	hi := 2 * l2
	mid1 := 0
	mid2 := 0
	L1 := 0
	L2 := 0
	R1 := 0
	R2 := 0
	for low <= hi {
		mid2 = (low + hi) / 2
		mid1 = l1 + l2 - mid2

		if mid1 == 0 {
			L1 = math.MinInt32
		} else {
			L1 = nums1[(mid1-1)/2]
		}

		if mid2 == 0 {
			L2 = math.MinInt32
		} else {
			L2 = nums2[(mid2-1)/2]
		}

		if mid1 == 2*l1 {
			R1 = math.MaxInt32
		} else {
			R1 = nums1[(mid1)/2]
		}

		if mid2 == 2*l2 {
			R2 = math.MaxInt32
		} else {
			R2 = nums2[(mid2)/2]
		}

		if L1 > R2 {
			low = mid2 + 1
			continue
		} else if L2 > R1 {
			hi = mid2 - 1
			continue
		} else {
			return (math.Max(float64(L1), float64(L2)) + math.Min(float64(R1), float64(R2))) / 2
		}
	}
	return -1
}
