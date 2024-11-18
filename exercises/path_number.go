package exercises

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	// Base case: if root is nil, return false
	if root == nil {
		return false
	}

	// If it's a leaf node, check if the remaining sum equals the node's value
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	// Recursively check left and right subtrees with reduced target sum
	newTarget := targetSum - root.Val
	return hasPathSum(root.Left, newTarget) || hasPathSum(root.Right, newTarget)
}
