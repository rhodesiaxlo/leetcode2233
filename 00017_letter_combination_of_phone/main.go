package main

import "fmt"

/**
获取第一个rune of string rune(str[0]) 更快
不要用 for .. range 然后复制
 */
func main() {
	fmt.Println(letterCombination("234"))
}

func letterCombination(digits string) []string {
	var ans []string

	// 注意边界情况
	if len(digits) == 0 {
		return ans
	}



	m := map[rune][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}
	var wfs func(s string, l string, )
	wfs = func(s string, l string) {
		if len(l) == 0 {
			ans = append(ans, s)
			return
		}

		//var c rune
		//for _, v := range l {
		//	c = v
		//	break
		//}

		arr, ok := m[rune(l[0])]
		if !ok {
			return
		}

		for _, v := range arr {
			wfs(s+v, l[1:])
		}
	}

	wfs("", digits)
	return ans
}
