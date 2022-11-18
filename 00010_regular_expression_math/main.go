package main

import "fmt"

func main() {
	fmt.Println(isMath("aa", "a*"))
}

/**
00056_merge_interval 案例
*/
func isMath(s string, p string) bool {
	szs, szp := len(s), len(p)
	var r [][]bool

	r = make([][]bool, szs+1)
	for i := range r {
		r[i] = make([]bool, szp+1)
	}

	s = " " + s
	p = " " + p
	r[0][0] = true

	for i := 0; i < szs+1; i++ { // 为什么这里是 0
		for j := 1; j < szp+1; j++ {  // 这里是 1 ？？？
			// if next letter is asterisk, ignore
			if j+1 < szp && string(p[j+1]) == "*" {
				continue
			}

			if i-1 >= 0 && string(p[j]) != "*" {
				r[i][j] = r[i-1][j-1] && (s[i] == p[j] || string(p[j]) == ".")
			} else if string(p[j]) == "*" {
				nonmath := j-2 >= 0 && r[i][j-2]
				onemath := i-1 >= 0 && r[i-1][j] && (s[i] == p[j-1] || string(p[j-1]) == ".")
				r[i][j] = nonmath || onemath
			}
		}
	}
	return r[szs][szp]
}
