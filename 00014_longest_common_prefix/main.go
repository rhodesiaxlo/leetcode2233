package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix3(strs))
}

/*
利用 sort 包合并同类项
ts:100%, ss:86%
*/
func longestCommonPrefix1(strs []string) string {
	sort.Strings(strs)
	head := strs[0]
	tail := strs[len(strs)-1]
	l := len(head)
	if l > len(tail) {
		l = len(tail)
	}
	var common []byte
	for i := 0; i < l; i++ {
		if head[i] != tail[i] {
			break
		}
		common = append(common, head[i])
	}
	return string(common)
}

/*
暴力解法 ts:51  ss:54
*/
func longestCommonPrefix2(strs []string) string {
	minL := math.MaxInt32
	for _, str := range strs {
		if len(str) < minL {
			minL = len(str)
		}
	}

	var common []byte
	for i := 0; i < minL; i++ {
		tmpByte := strs[0][i]
		for _, str := range strs {
			if str[i] != tmpByte {
				return string(common)
			}

		}
		common = append(common, tmpByte)
	}
	return string(common)
}

/*
分支算法, 用其他分支和第一个分支比较，然后保存公共， 然后取最小合集
ts:100%  ss: 54%
*/
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	ans := strs[0]
	l := len(strs)
	for i:=1;i< l; i++ {
		j := 0;
		l2 := len(strs[i])
		lans := len(ans)
		for j < l2 && j < lans && strs[i][j] == ans[j] {
			j++
		}

		if j >= lans {
			continue
		}

		ans = ans[:j]
	}
	return ans
}




