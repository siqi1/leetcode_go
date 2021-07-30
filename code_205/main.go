package main

import "fmt"

//给定两个字符串 s 和 t，判断它们是否是同构的。
//
//如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
//
//每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	_s := make(map[string]string)
	_t := make(map[string]string)
	for i := 0; i < len(s); i++ {
		sv, sh := _s[string(s[i])]
		tv, th := _t[string(t[i])]
		if (sh && sv != string(t[i])) || (th && tv != string(s[i])) {
			return false
		}
		_s[string(s[i])] = string(t[i])
		_t[string(t[i])] = string(s[i])

	}
	fmt.Println("_s:", _s)
	fmt.Println("_t", _t)
	return true
}

func main() {
	s := "papap"
	t := "titii"
	result := isIsomorphic(s, t)
	fmt.Println(result)
}
