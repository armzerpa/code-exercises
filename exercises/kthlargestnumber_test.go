package exercises

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestReArrengeArrayPartition(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		left     int
		right    int
		expected int
	}{
		{"Basic case", []int{3, 2, 1, 5, 6, 4}, 0, 5, 3},
		{"All equal", []int{1, 1, 1, 1, 1}, 0, 4, 4},
		{"Already sorted", []int{1, 2, 3, 4, 5}, 0, 4, 4},
		{"Reverse sorted", []int{5, 4, 3, 2, 1}, 0, 4, 0},
		{"Single element", []int{1}, 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reArrengeArrayPartition(tt.nums, tt.left, tt.right)
			if result != tt.expected {
				t.Errorf("reArrengeArrayPartition() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestQuickSelection(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := quickSelection(tt.nums, 0, len(tt.nums)-1, len(tt.nums)-tt.k)
			if result != tt.expected {
				t.Errorf("quickSelection() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFindKthLargest(t *testing.T) {
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
		{"Large random array", generateRandomArray(1000), 500, -1}, // -1 is a placeholder
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindKthLargest(tt.nums, tt.k)
			if tt.name == "Large random array" {
				// For large random array, we'll verify the result differently
				verifyLargeRandomArray(t, tt.nums, tt.k, result)
			} else if result != tt.expected {
				t.Errorf("FindKthLargest() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(10000) // Random integers between 0 and 9999
	}
	return arr
}

func verifyLargeRandomArray(t *testing.T, nums []int, k int, result int) {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedNums)))
	expected := sortedNums[k-1]
	if result != expected {
		t.Errorf("FindKthLargest() for large random array = %v, want %v", result, expected)
	}
}

func BenchmarkFindKthLargest(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
			nums := generateRandomArray(size)
			k := size / 2
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				FindKthLargest(nums, k)
			}
		})
	}
}
