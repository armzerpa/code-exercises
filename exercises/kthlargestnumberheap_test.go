package exercises

import (
	"fmt"
	"sort"
	"testing"
)

func TestFindKthLargestUsingHeap(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{"Basic case", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"All equal", []int{1, 1, 1, 1, 1}, 3, 1},
		{"Already sorted", []int{1, 2, 3, 4, 5}, 3, 3},
		{"Reverse sorted", []int{5, 4, 3, 2, 1}, 4, 2},
		{"Single element", []int{1}, 1, 1},
		{"Duplicate elements", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindKthLargestUsingHeap(tt.nums, tt.k)
			if result != tt.expected {
				t.Errorf("FindKthLargestUsingHeap() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFindKthLargestUsingHeapLarge(t *testing.T) {
	// Test with a large random array
	size := 100000
	nums := generateRandomArray(size)
	k := 1000

	result := FindKthLargestUsingHeap(nums, k)

	// Verify the result
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedNums)))
	expected := sortedNums[k-1]

	if result != expected {
		t.Errorf("FindKthLargestUsingHeap() for large array = %v, want %v", result, expected)
	}
}

func BenchmarkFindKthLargestUsingHeap(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
			nums := generateRandomArray(size)
			k := size / 2
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				FindKthLargestUsingHeap(nums, k)
			}
		})
	}
}
