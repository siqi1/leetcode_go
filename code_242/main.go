package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	fmt.Println('a')

	sindex := [26]int{}
	tindex := [26]int{}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d, %d\n", s[i], t[i])

		sindex[s[i]-'a']++
		tindex[t[i]-'a']++
	}

	fmt.Printf("%+v\n", sindex)
	fmt.Printf("%+v\n", tindex)

	return sindex == tindex
}

func main() {
	s := "anagram"
	t := "nagaram"

	isAnagram(s, t)
}
