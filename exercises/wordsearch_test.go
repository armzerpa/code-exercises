package exercises

import (
	"reflect"
	"testing"
)

func TestExist(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]byte
		word     string
		expected bool
	}{
		{
			name:     "1st case",
			input:    [][]byte{{'A', 'B', 'C'}, {'A', 'B', 'C'}, {'A', 'B', 'C'}},
			word:     "ABC",
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := Exist(tc.input, tc.word)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}
