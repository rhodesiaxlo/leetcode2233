package main

import "fmt"

func main() {
	i := 1994
	fmt.Println(intToRoman(i))
}

func intToRoman(i int) string {
	// 不可以使用 map , 因为 map 是无序的
	is := []int{1000, 500, 100, 50, 10, 5, 1}
	rs := []string{"M", "D", "C", "L", "X", "V", "I"}
	sz := len(is)
	// 2个的组合
	for i := 0; i < sz; i++ {
		for j := sz - 1; j > i; j-- {
			del := is[i] - is[j]
			if del == 900 || del == 400 || del == 90 || del == 40 || del == 9 || del == 4 {
				is = append(is, is[i]-is[j])
				rs = append(rs, rs[j]+rs[i])
			}
		}
	}

	esz := len(is)

	// 插入排序 zip sort
	// zip 插入排序 ！！！
	for i := 0; i < esz-1; i++ {
		bigi := i
		for j := i + 1; j < esz; j++ {
			if is[j] > is[bigi] {
				bigi = j
			}
		}
		if i != bigi {
			is[i], is[bigi] = is[bigi], is[i]
			rs[i], rs[bigi] = rs[bigi], rs[i]
		}
	}

	div := i
	ans := ""
	for div > 0 {
		for i, v := range is {
			if div >= v {
				m := div / v
				for m > 0 {
					ans += rs[i]
					m--
				}
				div = div % v
				break
			}
		}
	}
	return ans
}
