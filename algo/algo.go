package algo

import (
	"leetcode/leetcode/structure/list"
	"leetcode/leetcode/structure/stack"
	"math"
	"strings"
)

func LeetCode876MiddleOfLinkList(head *list.ListNode) *list.ListNode {
	if head == nil {
		return nil
	}

	s, f := head, head
	for s.Next != nil && f.Next != nil {
		s = s.Next
		f = f.Next
		if f.Next != nil {
			f = f.Next
		}
	}

	return s

	// time O(n)
	// space O(1)
}

func LeetCode26RemoveDumpFromSortedArray(nums []int, x int) int {
	sz := len(nums)
	for i := 0; i < sz; {
		tmpSz := len(nums)
		if i >= tmpSz {
			break
		}

		if i+1 < tmpSz && nums[i] == nums[i+1] {
			nums = append(nums[:i+1], nums[i+2:]...)

		} else {
			i++
		}
	}

	return len(nums)

	// time O(n)
	// space O(1)
	// ss 5%
}

func LeetCode26RemoveDumpFromSortedArray2(nums []int, x int) int {
	i, j := 0, 0
	sz := len(nums)
	for i = 0; i < sz; i++ {
		if nums[i] != nums[j] {
			j += 1
			nums[j] = nums[i]
		}
	}

	// 不需要这一步
	//nums = append([]int{}, nums[:j+1]...)
	return j + 1
}

/**
抽象一下， 删除重复大于k次的元素
*/
func LeetCode26RemoveDumpMorethanKFromSortedArray2(nums []int, x int, k int) int {
	// todo..
	return 0
}

func LeetCode27RemoveElement(nums []int, val int) int {
	i, j := 0, -1
	sz := len(nums)
	// 循环启动条件和动力
	for i = 0; i < sz; i++ {
		if nums[i] != val {
			j += 1
			nums[j] = nums[i]
		}
	}

	// 数组越界
	if j == -1 {
		nums = []int{}
	} else {
		nums = append([]int{}, nums[:j]...)
	}
	return j + 1
}

/** 迭代 **/
func LeetCode29divide2(dividend int, divisor int) int {
	// todo
	return 0
}

/** 递归 **/
func LeetCode29divide(dividend int, divisor int) int {
	// single special case
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	return div(dividend, divisor)
}

/**
00056_merge_interval 解法
*/
func div(dived int, d int) int {
	symbol := 1
	if (dived > 0 && d < 0) || (dived < 0 && d > 0) {
		symbol = -1
	}

	// 负数转整数可能会溢出
	if dived < 0 {
		dived = -dived
	}

	if d < 0 {
		d = -d
	}

	if dived < d {
		return 0
	}

	cnt := 1
	tmp := d
	for dived-tmp > tmp {
		cnt += cnt
		tmp += tmp
	}

	return symbol * (cnt + div(dived-tmp, d))
}

func LeetCode28IndexOfStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	if haystack == needle {
		return 0
	}

	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 {
		return -1
	}

	szH, szN := len(haystack), len(needle)
	for i := 0; i < szH-szN+1; i++ {
		if haystack[i] == needle[0] {
			j := i
			for j-i < szN && haystack[j] == needle[j-i] {
				j++
			}

			if j-i == szN {
				return i
			}
		}
	}
	return -1
}

/**
这个思路是有问题的， 因为 ( ) 不一定是连续出现的， 可能有间隔
 */
func LeetCode32longestValidParentheses(s string) int {
	// 溢出左边的 ) 和右边的 （
	s = strings.TrimLeft(s, ")")
	s = strings.TrimRight(s, "(")

	sz := len(s)
	if sz == 0 {
		return 0
	}
	maxN :=0
	st := stack.Stack{}
	for i := 0;i < sz;i ++ {
		st.Clear()
		for j:=i;j < sz; j++ {
			if s[j] == '(' {
				st.Push(rune(s[j]))
			}

			// if it is a math, push into the stack
			if s[j] == ')' && st.Peek() == '(' {
				st.Push(rune(s[j]))
			} else {
				// other wise, pairs break hero
				if maxN < st.Size() {
					maxN = st.Size()
				}
				// break from hero, jump to location
				i = j
				break
			}
		}
	}
	return maxN
}
