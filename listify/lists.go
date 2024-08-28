package listify

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res, surp := l1, 0
	for l1 != nil {
		temp := 0
		if l2 != nil {
			temp = l2.Val
			l2 = l2.Next
		}
		l1.Val = l1.Val + temp + surp
		surp = 0
		if l1.Val >= 10 {
			surp = 1
			l1.Val = l1.Val % 10
		}
		if l1.Next == nil {
			if surp != 0 {
				l1.Next = &ListNode{}
			}
			if l2 != nil {
				l1.Next = l2
				l2 = nil
			}
		}
		l1 = l1.Next
	}

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
	return res
}
