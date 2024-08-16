package exercises

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{
			name:       "Example 1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "Example 2",
			candidates: []int{2, 3, 5},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "Example 3",
			candidates: []int{2},
			target:     1,
			want:       [][]int{},
		},
		{
			name:       "Empty candidates",
			candidates: []int{},
			target:     5,
			want:       [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CombinationSum(tt.candidates, tt.target)
			if !compareSlicesIgnoreOrder(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to compare two slices of slices, ignoring order
func compareSlicesIgnoreOrder(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	aStr := make([]string, len(a))
	bStr := make([]string, len(b))

	for i, v := range a {
		sort.Ints(v)
		aStr[i] = sliceToString(v)
	}
	for i, v := range b {
		sort.Ints(v)
		bStr[i] = sliceToString(v)
	}

	sort.Strings(aStr)
	sort.Strings(bStr)

	return reflect.DeepEqual(aStr, bStr)
}

// Helper function to convert a slice of ints to a string
func sliceToString(slice []int) string {
	result := ""
	for _, v := range slice {
		result += string(rune(v + '0'))
	}
	return result
}
