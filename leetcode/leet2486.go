package leetcode

import "fmt"

func RunAppendCharacters() {
	s := "coaching"
	t := "cohing"
	res := appendCharacters(s, t)
	fmt.Println(res)
	fmt.Println(s, t)
}

func appendCharacters(s string, t string) int {
	r := 0
	if len(t) == 0 {
		return 0
	}
	for i := range s {
		if r >= len(t) {
			return 0
		}
		if s[i] == t[r] {
			r++
		}
	}
	if len(t) > r {
		return len(t) - r
	}
	return 0
}
