package exercises

import "testing"

// Helper function to create a tree from slice
func createTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		current := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != -1 {
			current.Left = &TreeNode{Val: values[i]}
			queue = append(queue, current.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != -1 {
			current.Right = &TreeNode{Val: values[i]}
			queue = append(queue, current.Right)
		}
		i++
	}

	return root
}

func TestIsSameTree(t *testing.T) {
	// Test case 1: Identical trees
	t.Run("Identical Trees", func(t *testing.T) {
		tree1 := createTree([]int{1, 2, 3})
		tree2 := createTree([]int{1, 2, 3})

		if !isSameTree(tree1, tree2) {
			t.Errorf("Trees should be identical")
		}
	})

	// Test case 2: Different trees with same structure
	t.Run("Different Value Trees", func(t *testing.T) {
		tree1 := createTree([]int{1, 2, 3})
		tree2 := createTree([]int{1, 2, 4})

		if isSameTree(tree1, tree2) {
			t.Errorf("Trees should not be identical")
		}
	})

	// Test case 3: Empty trees
	t.Run("Empty Trees", func(t *testing.T) {
		if !isSameTree(nil, nil) {
			t.Errorf("Two nil trees should be considered the same")
		}
	})

	// Test case 4: One empty, one non-empty tree
	t.Run("One Empty Tree", func(t *testing.T) {
		tree := createTree([]int{1, 2, 3})

		if isSameTree(tree, nil) || isSameTree(nil, tree) {
			t.Errorf("One empty and one non-empty tree should not be the same")
		}
	})

	// Test case 5: Different structures
	t.Run("Different Structures", func(t *testing.T) {
		tree1 := createTree([]int{1, 2, 3})
		tree2 := createTree([]int{1, 2, -1, 3})

		if isSameTree(tree1, tree2) {
			t.Errorf("Trees with different structures should not be the same")
		}
	})

	// Test case 6: Complex tree comparison
	t.Run("Complex Tree Comparison", func(t *testing.T) {
		tree1 := createTree([]int{4, 2, 7, 1, 3, 6, 9})
		tree2 := createTree([]int{4, 2, 7, 1, 3, 6, 9})
		tree3 := createTree([]int{4, 2, 7, 1, 3, 6, 10})

		if !isSameTree(tree1, tree2) {
			t.Errorf("Identical complex trees should be the same")
		}

		if isSameTree(tree1, tree3) {
			t.Errorf("Trees with different values should not be the same")
		}
	})
}
