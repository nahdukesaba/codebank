package leetcode

import (
	"fmt"
	"strconv"
)

func RunNumSteps() {
	input := "1101"
	res := numSteps(input)
	fmt.Println(res)
}

func numSteps(s string) int {
	i, _ := strconv.ParseInt(s, 2, 64)
	n := 0
	for i != 1 {
		if i%2 == 0 {
			i /= 2
		} else {
			i++
		}
		n++
	}
	return n
}
