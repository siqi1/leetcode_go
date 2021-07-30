package main

import "fmt"

func isValid(s string) bool {
	rule := make(map[byte]byte)
	rule[')'] = '('
	rule[']'] = '['
	rule['}'] = '{'
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		//存在
		ru, ex := rule[s[i]]
		if len(stack) > 0 && ex {
			//弹栈
			_stack := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if _stack != ru {
				return false
			}
		} else {
			//压栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func main() {
	s := "]]"
	result := isValid(s)
	fmt.Println("result:", result)
}
