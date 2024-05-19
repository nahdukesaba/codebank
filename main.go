package main

import (
	"fmt"
	"test-code/stringify"
)

func main() {
	a := "18 42"
	b := "asdfqwer"
	fmt.Println(stringify.InterlacedString(a, b))
}
