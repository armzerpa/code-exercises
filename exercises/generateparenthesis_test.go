package exercises

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	// Test case for n = 0
	//result0 := generateParenthesis(0)
	/*if len(result0) != 0 {
		t.Errorf("Expected empty slice for n = 0, got %v", result0)
	}*/

	// Test case for n = 1
	result1 := generateParenthesis(1)
	expected1 := []string{"()"}
	sortAndCompare(t, result1, expected1)

	// Test case for n = 3
	result3 := generateParenthesis(3)
	expected3 := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	sortAndCompare(t, result3, expected3)

	// Validate that all generated strings are valid
	for _, combination := range result3 {
		if !isValidParenthesesCombination(combination) {
			t.Errorf("Invalid parentheses combination: %s", combination)
		}
	}
}

// Helper function to sort and compare slices
func sortAndCompare(t *testing.T, actual, expected []string) {
	sort.Strings(actual)
	sort.Strings(expected)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

// Helper function to validate parentheses combination
func isValidParenthesesCombination(s string) bool {
	balance := 0
	for _, c := range s {
		if c == '(' {
			balance++
		} else {
			balance--
		}

		// At any point, balance should not be negative
		if balance < 0 {
			return false
		}
	}

	// At the end, balance should be zero
	return balance == 0
}

// Benchmark the function
func BenchmarkGenerateParenthesis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesis(3)
	}
}
