package exercises

import (
	"testing"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{3, 0, 1},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 1},
			expected: 2,
		},
		{
			name:     "Example 3",
			nums:     []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			expected: 8,
		},
		{
			name:     "Single element",
			nums:     []int{0},
			expected: 1,
		},
		{
			name:     "Missing first number",
			nums:     []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "Missing last number",
			nums:     []int{0, 1, 2},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MissingNumber(tt.nums)
			if result != tt.expected {
				t.Errorf("missingNumber(%v) = %d; want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
