package main

//  * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	if root.Val == root.Right.Val+root.Left.Val {
		return true
	} else {
		return false
	}
}
