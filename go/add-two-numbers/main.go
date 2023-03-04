package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c, tmp := 0, new(ListNode)
	ans := tmp
	for l1 != nil || l2 != nil || c != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}
		n := n1 + n2 + c
		c = n / 10
		fmt.Println(n % 10)
		ans.Next = &ListNode{n % 10, nil}
		ans = ans.Next
	}
	return tmp.Next
}
