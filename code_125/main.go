package main

import (
	"fmt"
	"strings"
)

func check(str byte) bool {
	if (str >= 'a' && str <= 'z') || (str >= '0' && str <= '9') {
		return true
	}
	return false
}

//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//说明：本题中，我们将空字符串定义为有效的回文串。
func isPalindrome(s string) bool {
	//字符串小写
	_s := strings.ToLower(s)

	i, j := 0, len(_s)-1

	for i < j {
		if !check(_s[i]) {
			i++
			continue
		} else if !check(_s[j]) {
			j--
			continue
		}
		if _s[i] != _s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	s := "A？》《》《《 man, a plan, a canal: Panam a"
	result := isPalindrome(s)
	fmt.Println("result:", result)
}
