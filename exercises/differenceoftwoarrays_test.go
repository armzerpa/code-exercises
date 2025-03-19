package exercises

import (
	"testing"
)

func TestFindDifference(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{"case 1", []int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findDifference(tt.nums1, tt.nums2)
			if !compareSlicesIgnoreOrder(result, tt.expected) {
				t.Errorf("findDifference() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFindDifference2(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{"case 1", []int{1, 2, 3}, []int{2, 4, 6}, [][]int{{1, 3}, {4, 6}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findDifference2(tt.nums1, tt.nums2)
			if !compareSlicesIgnoreOrder(result, tt.expected) {
				t.Errorf("findDifference() = %v, want %v", result, tt.expected)
			}
		})
	}
}
