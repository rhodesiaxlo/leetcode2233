package main

import "fmt"

func main() {
	s := "abcdefghijk"
	fmt.Println(zigzagConvert(s, 3))
}

func zigzagConvert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var rows [][]byte
	i := 0
	for i := 0; i < numRows; i++ {
		rows = append(rows, []byte(""))
	}

	godown := 1
	for _, c := range s {
		rows[i] = append(rows[i], byte(c))
		i += godown // 第一次作为初始条件，已经走过了， 所以不会重走， start
		if i == -1 {
			i = 1
			godown = 1
		} else if i == numRows {
			i = numRows - 2
			godown = -1
		}
	}

	ret := ""
	for i := 0; i < numRows; i++ {
		ret += string(rows[i])
	}
	return ret
}
