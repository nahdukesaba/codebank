package main

import "test-code/query"

//func main() {
//	req1 := &listify.ListNode{
//		Val: 0,
//	}
//	req2 := &listify.ListNode{
//		Val: 2,
//		Next: &listify.ListNode{
//			Val: 7,
//			Next: &listify.ListNode{
//				Val: 8,
//			},
//		},
//	}
//	resList := listify.AddTwoNumbers(req1, req2)
//	fmt.Println("test", resList)
//
//	a := "18 42"
//	b := "asdfqwer"
//	fmt.Println(stringify.InterlacedString(a, b))
//
//	exp := "2 + (1 * 3)"
//	res := stringify.SimpleCalculateWithSpaces(exp)
//	fmt.Println(res)
//
//	exp2 := "2+1*3"
//	res2 := stringify.SimpleCalculate(exp2)
//	fmt.Println(res2)
//
//	//fs := token.NewFileSet()
//	//tr, _ := parser.ParseExpr("(3-1) * 5")
//	//ast.Print(fs, tr)
//	//tr.Pos()
//
//	expression := "2 + 1 * 3"
//	result, err := EvaluateExpression(expression)
//	if err != nil {
//		fmt.Println("Error evaluating expression:", err)
//		return
//	}
//	fmt.Printf("Result: %d\n", result)
//
//	req := []int{2}
//	res = slicify.RemoveElement(req, 3)
//	fmt.Println(res)
//
//	req3 := []int{1, 2, 3, 4}
//	res3 := slicify.ProductExceptSelf(req3)
//	fmt.Println(res3)
//
//	leetcode.RunLengthOfLongestSubstring()
//	leetcode.RunAppendCharacters()
//	leetcode.RunNumSteps()
//	leetcode.RunHeightChecker()
//}

func main() {
	query.TestQuery()
}
