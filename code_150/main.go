package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, v := range tokens {
		switch v {
		case "+", "-", "*", "/":
			n1, n2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			var calc int
			switch v {
			case "+":
				calc = n1 + n2
			case "-":
				calc = n1 - n2
			case "*":
				calc = n1 * n2
			case "/":
				calc = n1 / n2
			}
			stack = append(stack, calc)
		default:
			num, _ := strconv.Atoi(v)
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	result := evalRPN(tokens)
	fmt.Println("result :", result)
}
