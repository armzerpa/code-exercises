package exercises

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Basic case with zeros",
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "Array with no zeros",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Array with all zeros",
			input:    []int{0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Mixed zeros and non-zeros at different positions",
			input:    []int{0, 0, 1, 0, 3, 0, 12, 0},
			expected: []int{1, 3, 12, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a copy of the input to modify
			nums := make([]int, len(tc.input))
			copy(nums, tc.input)

			// Call the function
			moveZeroes(nums)

			// Check if the result matches the expected output
			if !reflect.DeepEqual(nums, tc.expected) {
				t.Errorf("Failed for case %s: got %v, want %v",
					tc.name, nums, tc.expected)
			}
		})
	}
}
