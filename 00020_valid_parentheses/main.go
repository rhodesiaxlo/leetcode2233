package main

import "fmt"

/*
1. rune string bytes 的区别
2. struct 指针方法和非指针方法
 */
func main() {
	s := "(]"
	fmt.Println(isValid(s))
}

/**
第一版， 有类型转化  ss:45%
 */
func isValid(s string) bool {
	ss := stack1{}
	for _, v := range s {
		tmp := string(v)
		if tmp == "[" || tmp == "{" || tmp == "(" {
			ss.push(tmp)
			continue
		}

		p, err := ss.pop()
		if err != nil {
			return false
		}
		if tmp == ")" && p != "(" {
			return false
		}

		if tmp == "}" && p != "{" {
			return false
		}

		if tmp == "]" && p != "[" {
			return false
		}

	}
	return ss.isEmpty()
}

type stack1 struct {
	arr2 []string
}

func (s *stack1) isEmpty() bool {
	return len(s.arr2) == 0
}

func (s *stack1) push(str string) {
	s.arr2 = append(s.arr2, str)
}

func (s *stack1) pop() (string, error) {
	sz := len(s.arr2)
	if sz == 0 {
		return "", fmt.Errorf("error")
	}
	tmp := s.arr2[sz-1]
	s.arr2 = append(s.arr2[:sz-1], s.arr2[sz:]...)
	return tmp, nil
}

/**
第二版 没有类型转化 ss:100%
 */
func isValid2(s string) bool {
	ss := stack2{}
	for _, v := range s {
		if v == '[' || v == '{' || v == '(' {
			ss.Push(v)
			continue
		}

		p, err := ss.Pop()
		if err != nil {
			return false
		}
		if v == ')' && p != '(' {
			return false
		}

		if v == '}' && p != '{' {
			return false
		}

		if v == ']' && p != '[' {
			return false
		}

	}
	return ss.IsEmpty()
	
}

type stack2 struct {
	arr2 []rune // 类型转换很消耗性能
}

func (s2 *stack2) IsEmpty() bool {
	return len(s2.arr2) == 0
}

func (s2 *stack2) Push(r rune)  {
	s2.arr2 = append(s2.arr2, r)
}

func (s2 *stack2) Pop() (rune, error) {
	if len(s2.arr2) == 0 {
		return ' ', fmt.Errorf("error")
	}
	sz := len(s2.arr2)
	ans := s2.arr2[sz-1]
	s2.arr2 = append(s2.arr2[:sz-1], s2.arr2[sz:]...)
	return ans,nil
}




