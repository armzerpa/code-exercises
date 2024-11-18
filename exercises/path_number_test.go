package exercises

import (
	"testing"
)

// helper function to create test trees
func createTestTree(values []int, nullIndices map[int]bool) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(values))

	// Create nodes
	for i, val := range values {
		if !nullIndices[i] {
			nodes[i] = &TreeNode{Val: val}
		}
	}

	// Connect nodes
	for i := 0; i < len(values); i++ {
		if nodes[i] != nil {
			leftIdx := 2*i + 1
			rightIdx := 2*i + 2

			if leftIdx < len(nodes) {
				nodes[i].Left = nodes[leftIdx]
			}
			if rightIdx < len(nodes) {
				nodes[i].Right = nodes[rightIdx]
			}
		}
	}

	return nodes[0]
}

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		values    []int
		nullNodes map[int]bool
		targetSum int
		want      bool
	}{
		{
			name:      "Empty tree",
			values:    []int{},
			nullNodes: map[int]bool{},
			targetSum: 0,
			want:      false,
		},
		{
			name:      "Single node equals target",
			values:    []int{1},
			nullNodes: map[int]bool{},
			targetSum: 1,
			want:      true,
		},
		{
			name:      "Single node not equals target",
			values:    []int{1},
			nullNodes: map[int]bool{},
			targetSum: 2,
			want:      false,
		},
		{
			name:      "Standard tree with valid path",
			values:    []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1},
			nullNodes: map[int]bool{4: true, 9: true, 10: true, 11: true},
			targetSum: 22,
			want:      true,
		},
		{
			name:      "Standard tree without valid path",
			values:    []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1},
			nullNodes: map[int]bool{4: true, 9: true, 10: true, 11: true},
			targetSum: 100,
			want:      false,
		},
		{
			name:      "Negative values with valid path",
			values:    []int{-2, -3, 4, -1, -5},
			nullNodes: map[int]bool{3: true},
			targetSum: -10,
			want:      true,
		},
		{
			name:      "All negative values without valid path",
			values:    []int{-2, -3, -4},
			nullNodes: map[int]bool{},
			targetSum: 0,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := createTestTree(tt.values, tt.nullNodes)
			got := hasPathSum(root, tt.targetSum)
			if got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Additional test cases for edge cases
func TestHasPathSumEdgeCases(t *testing.T) {
	t.Run("Nil root", func(t *testing.T) {
		if hasPathSum(nil, 0) {
			t.Error("Expected false for nil root")
		}
	})

	t.Run("Zero sum with zero value", func(t *testing.T) {
		root := &TreeNode{Val: 0}
		if !hasPathSum(root, 0) {
			t.Error("Expected true for zero sum with single zero node")
		}
	})

	t.Run("Deep tree with single valid path", func(t *testing.T) {
		root := &TreeNode{Val: 1}
		current := root
		sum := 1
		// Create a deep right-skewed tree
		for i := 2; i <= 5; i++ {
			current.Right = &TreeNode{Val: i}
			current = current.Right
			sum += i
		}
		if !hasPathSum(root, sum) {
			t.Errorf("Expected true for sum %d in deep tree", sum)
		}
	})
}
