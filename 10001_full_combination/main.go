package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	x := []int{1, 2, 3, 4, 4}
	com := combination(x)
	com4 := combination4(x)
	fmt.Print(com4)
	fmt.Printf("n=%d, combination number = %d\n", len(x), len(com))

	m := map[int]string{1000: "M", 500: "D", 100: "C", 50: "L", 10: "X", 5: "V", 1: "I"}
	fmt.Println(m)
}

/**
思考一下如何去重
*/
func combination(x []int) []string {
	com := make([]string, 200)
	var dfs func(s string, seeds []int)
	n := 0

	dfs = func(s string, seeds []int) {
		if len(seeds) == 0 {
			com[n] = s
			n++
			return
		}

		for i, v := range seeds {
			// 注意这里如果 src 有容量会修改原来的， 有坑
			var tmp []int
			tmp = append(tmp, seeds[:i]...)
			tmp = append(tmp, seeds[i+1:]...)
			dfs(s+strconv.Itoa(v), tmp)
		}
	}

	dfs("", x)
	return com
}

/**
backtracking, n 个 slot, 每个 slot 中有 n 个可能， 判断组合是否重复
*/
func combination3(x []int) []string {
	// 先全排

	// 再去重

	return nil
}

/**
字典序求解，我们知道 长度为n, 元素互补重复的排列组合为 n!, 所以，我们按照一定的顺序把剩余的找出来就好了
*/
func combination4(x []int) []string {
	// 首先排序， 升序排列，此时数组产生的排列是最小的
	sort.Ints(x)
	var ans []string

	for {
		ans = append(ans, arrToString(x))
		x = nextPermutation(x)
		if x == nil {
			break
		}
	}
	return ans
}

/**
这个算法和很强的学习意义
1. 题目就很有意义，给出每次移动一步， 这一步是最小的改变，来产生下一个排列
2. 在解题的时候，先排序
3. 排序找到第一个逆序对
4. 然后根据排序的规则， 找到第一个比 i 大的 j , 这个 j 则一定满组 j 是最接近 i 且比 i 大的元素
5. 最后因为高位大了一点点， 所以低位就必须重新排序达到一个最小值， 因为定义的原因， 逆序就好了
6. 如果元素中有重复元素 （精髓)

判断索引在取值钱
判断模数在整除前
*/
func nextPermutation(x []int) []int {
	sz := len(x)
	i := sz - 2
	// 按照 permutation 的原则， 最小移动一位
	for i >= 0 && x[i] >= x[i+1] { // !!! 索引判定必须在元素的前面
		i--
	}

	// 已经是最大值了，找不到下一个组合
	if i < 0 {
		return nil
	}

	// 找到第一个比 x[i] 大的元素，按照定义，第一个碰到的就是最接近 x[i] 且比 x[i] 大的元素
	j := sz - 1
	for j >= 0 && x[j] <= x[i] {
		j--
	}
	x[i], x[j] = x[j], x[i]

	// 高位给了一个大一点点的数之后， 剩下的可以恢复到最小的状态
	// 这里是一个排序算法
	l, r := i+1, sz-1
	for l < r {
		x[l], x[r] = x[r], x[l]
		l++
		r--
	}
	return x
}

/**
如果数组是重复的怎么办, 去重
*/
func combination6(x []int) []string {
	return []string{""}
}

func arrToString(i []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(i), " ", "", -1), "[]")
}
