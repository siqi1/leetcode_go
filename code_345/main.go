package main

import "fmt"

// 编写一个函数，以字符串作为输入，反转该字符串中的元音字母。
//func reverseVowels(s string) string {
//	strArr := []byte(s)
//	index := 0
//	for k,v := range strArr {
//		if v == 'a' || v == 'o' || v == 'e' || v == 'i' || v == 'u' {
//			if index != 0 {
//				a := strArr[k]
//				b := strArr[index]
//				strArr[index] = a
//				strArr[k] = b
//			}
//			index = k
//			fmt.Println("index:",index)
//		}
//	}
//	return string(strArr)
//}

func check(s byte) bool {
	switch s {
	case 'a', 'o', 'e', 'i', 'u', 'A', 'O', 'E', 'I', 'U':
		return true
	default:
	}
	return false
}

func reverseVowels(s string) string {
	strArr := []byte(s)
	j, k := 0, len(strArr)-1
	for j < k {
		if !check(strArr[j]) {
			j++
			continue
		}
		if !check(strArr[k]) {
			k--
			continue
		}
		a := strArr[j]
		b := strArr[k]
		strArr[j] = b
		strArr[k] = a
		j++
		k--
	}
	return string(strArr)
}

func main() {
	s := "aA"
	result := reverseVowels(s)
	fmt.Println("result:", result)
}
