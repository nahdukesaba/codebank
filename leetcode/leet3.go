package leetcode

import "fmt"

func RunLengthOfLongestSubstring() {
	input := "pwwkew"
	res := LengthOfLongestSubstring(input)
	fmt.Println(res)
}

func LengthOfLongestSubstring(s string) int {
	mapString := map[string]int{}
	l := 0
	//m := 0
	for _, v := range s {
		//m = max(m, len(mapString))
		val := string(v)
		mapString[val]++

		for mapString[val] > 1 {
			val2 := string(s[l])
			mapString[val2]--
			if mapString[val2] == 0 {
				delete(mapString, val2)
			}
			l++
		}
	}

	//return max(m, len(mapString))
	return len(mapString)
}
