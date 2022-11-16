package main

import (
	"fmt"
)

/**
1. 思考一下等效替换, 每次下一对括号，为什么就能覆盖括号分开放的情况 ？？
2. 递归还是迭代
 */
func main() {
	fmt.Println(generateParenthesis2(3))
}

func generateParenthesis(n int) []string {
	ans := [][]string{[]string{""}, []string{"()"}}
	if n == 1 {
		return ans[1]
	}

	for i := 2; i <= n; i++ {
		var tmp []string
		visited := make(map[string]bool)
		for _, v := range ans[i-1] {
			for j := 0; j < 2*(i-1)+1; j++ {
				news := v[:j] + "()" + v[j:]
				if y, ok := visited[news]; !ok || !y {
					tmp = append(tmp, news)
					visited[news] = true
				}
			}
		}
		ans = append(ans, tmp)
	}

	return ans[n]
}

/**

 */
func generateParenthesis2(n int) []string {
	ans := make([]string, 0)
	current := make([]byte, n*2)
	rec22(&ans, n, 0, 0, current)
	return ans
}

func rec22(ans *[]string, n int, left int, right int, current []byte) {
	if left+right == n*2 {
		*ans = append(*ans, string(current))
		return
	}

	if left < n {
		current[left+right] = '('
		rec22(ans, n, left+1, right, current)
	}

	if right < left {
		current[left+right] = ')'
		rec22(ans, n, left, right+1, current)
	}
}