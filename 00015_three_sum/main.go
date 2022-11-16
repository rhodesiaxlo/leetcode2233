package main

/**
1. bs 有局限性， 只能返回第一个元素， 如果有重复元素的话不适用
2. 分成2段会有一个印象就是，如果第一段为 Length = 0， 循环跑不起来
 */
import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//sum := []int{0, 0, 0}
	//sum2 := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6} // 大量重复元素
	sum3 := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	//fmt.Println(threeSum(sum))
	fmt.Println(threeSum(sum3)) // 重复元素 bs 有局限性
}

/**
暴力破解
*/
func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	sz := len(nums)

	if nums[0] == 0 && nums[sz-1] == 0 {
		return [][]int{[]int{0, 0, 0}}
	}

	var l, r []int
	// split

	for i := 0; i < sz-1; i++ {
		if nums[i] == 0 { // 保证初始循环可以走
			l = nums[:i+1]
			r = nums[i+1:]
			break

		} else if nums[i+1] > 0 {
			l = nums[:i+1]
			r = nums[i+1:]
			break
		}
	}

	// 去重
	var visited = make(map[string]bool)
	szl, szr := len(l), len(r)
	for i := 0; i < szl; i++ {
		for j := 0; j < szr; j++ {
			remain := 0 - l[i] - r[j]
			if remain >= 0 {
				if index := bs(remain, r, j); index != -1 && index != j {
					f, m, la := getOrderSeq(l[i], r[j], r[index])
					offset := getOffset(f, m, la)
					if visited[offset] == false {
						ans = append(ans, []int{l[i], r[j], r[index]})
						visited[offset] = true
					}
				}
			} else {
				if index := bs(remain, l, i); index != -1 && index != i {
					f, m, la := getOrderSeq(l[i], r[j], l[index])
					offset := getOffset(f, m, la)
					if visited[offset] == false {
						ans = append(ans, []int{l[i], r[j], l[index]})
						visited[offset] = true
					}
				}
			}
		}
	}

	// slice 去重

	return ans
}

func bs(t int, src []int, e int) int {
	sz := len(src)
	l, r := 0, sz-1
	for l <= r {
		mid := (r-l)/2 + l
		if src[mid] == t {
			if mid != e {
				return mid
			} else {
				l = mid + 1
			}
		}

		if src[mid] > t {
			r = mid - 1
		}

		if src[mid] < t {
			l = mid + 1
		}
	}

	return -1
}

func removeDump(src [][]int) {
	sz := len(src)
	for i := 0; i < sz-1; {
		if i+1 < len(src) && src[i][0] == src[i+1][0] {
			if (src[i][1] == src[i+1][1] && src[i][2] == src[i+1][2]) ||
				(src[i][1] == src[i+1][2] && src[i][2] == src[i+1][1]) {
				src = append(src[:i+1], src[i+2:]...)
			} else {
				i++
			}
		}
	}
}

func getOrderSeq(i, j, k int) (int, int, int) {
	src := []int{i, j, k}
	sort.Ints(src)
	return src[0], src[1], src[2]
}

func getOffset(i, j, k int) string {
	return strconv.Itoa(i*2 + j*13 + k*23)
}
