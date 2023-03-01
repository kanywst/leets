package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) (ans bool) {
	ansLeft, ansRight := false, false
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val == q.Val {
		ansLeft, ansRight = isSameTree(p.Left, q.Left), isSameTree(p.Right, q.Right)
	} else {
		return false
	}
	if ansLeft && ansRight {
		ans = true
	} else {
		ans = false
	}
	return
}
