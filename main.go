package main

import (
	"fmt"
	"test-code/stringify"
)

func main() {
	a := "18 42"
	b := "asdfqwer"
	fmt.Println(stringify.InterlacedString(a, b))

	exp := "2 + 1 * 3"
	res := stringify.SimpleCalculateWithSpaces(exp)
	fmt.Println(res)

	exp2 := "2+1*3"
	res2 := stringify.SimpleCalculate(exp2)
	fmt.Println(res2)
}
