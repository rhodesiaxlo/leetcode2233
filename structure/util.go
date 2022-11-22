package structure

import (
	"fmt"
	"leetcode/leetcode/structure/tree"
	"math"
	"strconv"
	"strings"
)

/**
nums 数组，代表一个完全数, 必须基于 1
*/
func PrintComTree(nums []int) {
	arr, _ := _div(nums, 0)
	for _, v := range arr {
		tmp := strings.Builder{}
		for _, v2 := range v {
			if v2 == "" {
				tmp.WriteString(" ")
			} else {
				tmp.WriteString(v2)
			}
		}
		fmt.Println(tmp.String())
	}
}

func _div(nums []int, level int) ([][]string, int) {
	sz := len(nums)
	eleTotal := sz - 1
	eleCount := 0
	upperLimit := int(math.Pow(float64(2), float64(level+1)) - 1)
	maxEle := int(math.Pow(float64(2), float64(level)))
	if upperLimit < eleTotal {
		eleCount = maxEle
	} else {
		eleCount = eleTotal - int(math.Pow(float64(2), float64(level))) + 1
	}

	var arr [][]string
	var maxWidth int
	if upperLimit < eleTotal {
		arr, maxWidth = _div(nums, level+1) // DP
	} else {
		// 最后一层, 构造 arr, 和 maxWidth 返回回去
		arr = make([][]string, level+1)
		maxWidth = 2*maxEle - 1
	}

	// 填充
	arr[level] = make([]string, maxWidth)
	eleWidth := (maxWidth - maxEle + 1) / maxEle
	mid := eleWidth/2 + 1
	for i, j := 0, 0; i < eleCount; i++ {
		arr[level][j+mid-1] = strconv.Itoa(nums[int(math.Pow(float64(2), float64(level)))+i])
		j += eleWidth + 1 // 这个语法有点意思
	}
	return arr, maxWidth
}


func dive(nodes []*tree.TreeNode, level int) ([][]string, int) {
	count := len(nodes)
	nextLevelExists := false

	// Find if there is another level
	// and construct next level filling up missing nodes as `nil`
	for i := 0; i < count; i++ {
		node := nodes[i]
		if node == nil {
			nodes = append(nodes, nil, nil)
		} else {
			nodes = append(nodes, node.L, node.R)
		}

		nextLevelExists = nextLevelExists || (node != nil && (node.L != nil || node.R != nil))
	}

	// Initialize array if we got to the bottom of the tree
	// otherwise, retrieve array from another recursive call
	var arr [][]string
	var width int
	if nextLevelExists {
		arr, width = dive(nodes[count:], level + 1)
	} else {
		arr = make([][]string, level + 1)
		width = 2 * count - 1
	}

	arr[level] = make([]string, width)
	elementWidth := (width - count + 1) / count

	// Processing current level place node.Val value into correct position in array
	for i, j := 0, 0; i < count; i++ {
		mid := elementWidth / 2 + 1

		if nodes[i] != nil {
			arr[level][j + mid - 1] = strconv.Itoa(nodes[i].V)
		}
		j += elementWidth + 1
	}
	return arr, width
}

func getArrForNonCompleteTree(root *tree.TreeNode) [][]string {
	arr, _ := dive([]*tree.TreeNode{root}, 0)
	return arr
}

func PrintNonCompleteTree(root *tree.TreeNode)  {
	arr := getArrForNonCompleteTree(root)
	for _, v := range arr {
		tmp := strings.Builder{}
		for _, v2 := range v {
			if v2 == "" {
				tmp.WriteString(" ")
			} else {
				tmp.WriteString(v2)
			}
		}
		fmt.Println(tmp.String())
	}
	//fmt.Println(arr)
}