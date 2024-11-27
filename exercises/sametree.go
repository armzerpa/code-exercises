package exercises

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// If both trees are nil, they are the same
	if p == nil && q == nil {
		return true
	}

	// If one tree is nil and the other is not, they are different
	if p == nil || q == nil {
		return false
	}

	// Check if current node values are the same
	if p.Val != q.Val {
		return false
	}

	// Recursively check left and right subtrees
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
