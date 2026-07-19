package p3997

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countDominantNodes(root *TreeNode) int {
	var res int
	var f func(node *TreeNode) int

	f = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		w := node.Val
		w = max(w, f(node.Left))
		w = max(w, f(node.Right))

		if w == node.Val {
			res++
		}
		return w
	}

	f(root)

	return res
}
