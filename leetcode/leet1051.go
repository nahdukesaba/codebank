package leetcode

import (
	"fmt"
	"sort"
)

func RunHeightChecker() {
	input := []int{5, 1, 2, 3, 4}
	res := heightChecker(input)
	fmt.Println(res)
	fmt.Println("aaa")
}

func heightChecker(heights []int) int {
	hs := heights
	fmt.Println(heights)
	fmt.Println(hs)
	sort.Ints(heights)
	fmt.Println(heights)
	fmt.Println(hs)
	res := 0
	for i, v := range hs {
		if v != heights[i] {
			res++
		}
	}
	return res
}
