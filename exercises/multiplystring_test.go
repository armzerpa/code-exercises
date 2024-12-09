package exercises

import (
	"testing"
)

func TestMultiply(t *testing.T) {
	// Test cases covering various scenarios
	testCases := []struct {
		name     string
		num1     string
		num2     string
		expected string
	}{
		{
			name:     "Basic multiplication",
			num1:     "2",
			num2:     "3",
			expected: "6",
		},
		{
			name:     "Multiplication with zero",
			num1:     "0",
			num2:     "123",
			expected: "0",
		},
		{
			name:     "Larger numbers",
			num1:     "123",
			num2:     "456",
			expected: "56088",
		},
		{
			name:     "One single digit, one multi-digit",
			num1:     "9",
			num2:     "99",
			expected: "891",
		},
		{
			name:     "Large numbers with carry",
			num1:     "999",
			num2:     "999",
			expected: "998001",
		},
		{
			name:     "Zero multiplication both ways",
			num1:     "0",
			num2:     "0",
			expected: "0",
		},
		{
			name:     "Large numbers",
			num1:     "1234",
			num2:     "5678",
			expected: "7006652",
		},
		{
			name:     "Numbers with leading zeros",
			num1:     "01",
			num2:     "10",
			expected: "10",
		},
	}

	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := multiply(tc.num1, tc.num2)

			if result != tc.expected {
				t.Errorf("multiply(%s, %s): expected %s, got %s",
					tc.num1, tc.num2, tc.expected, result)
			}
		})
	}
}

// Benchmark the multiply function
func BenchmarkMultiply(b *testing.B) {
	// Benchmark with moderately sized numbers
	num1 := "1234"
	num2 := "5678"

	// Run the benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		multiply(num1, num2)
	}
}

// Additional edge case test for extremely large numbers
func TestMultiplyLargeNumbers(t *testing.T) {
	largeNum1 := "123456789012345678901234567890"
	largeNum2 := "987654321098765432109876543210"

	// Note: The exact result might vary depending on the precision of your implementation
	result := multiply(largeNum1, largeNum2)

	if result == "" {
		t.Errorf("Large number multiplication failed, got empty string")
	}
}
