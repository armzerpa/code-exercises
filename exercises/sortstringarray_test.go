package exercises

import (
	"reflect"
	"testing"
)

func TestSortPriorityTimestamp(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "Sort by priority and timestamp",
			input:    []string{"2:1500", "1:1000", "3:2000", "1:1200", "5:3000"},
			expected: []string{"1:1000", "1:1200", "2:1500", "3:2000", "5:3000"},
		},
		{
			name:     "Single priority elements",
			input:    []string{"1:100", "1:50", "1:200"},
			expected: []string{"1:50", "1:100", "1:200"},
		},
		{
			name:     "Mixed priorities",
			input:    []string{"5:3000", "1:100", "3:2000", "2:1500"},
			expected: []string{"1:100", "2:1500", "3:2000", "5:3000"},
		},
		{
			name:     "Empty array",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := SortPriorityTimestamp(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

// Additional test for error handling
func TestSortPriorityTimestampInvalidFormat(t *testing.T) {
	input := []string{"2:1500", "invalid", "1:1000"}
	result := SortPriorityTimestamp(input)

	// Ensure function doesn't crash with invalid input
	if len(result) != len(input) {
		t.Errorf("Expected same length array, got %d instead of %d", len(result), len(input))
	}
}
