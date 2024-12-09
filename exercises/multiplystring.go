package exercises

import "strings"

func multiply(num1 string, num2 string) string {
	// Handle special cases where either number is "0"
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// Reverse the strings to make calculations easier
	n1 := reverseString(num1)
	n2 := reverseString(num2)

	// Initialize result array to store multiplication
	result := make([]int, len(n1)+len(n2))

	// Perform multiplication digit by digit
	for i := 0; i < len(n1); i++ {
		for j := 0; j < len(n2); j++ {
			// Convert characters to integers
			digit1 := int(n1[i] - '0')
			digit2 := int(n2[j] - '0')

			// Multiply digits and add to result
			product := digit1 * digit2
			result[i+j] += product
			result[i+j+1] += result[i+j] / 10
			result[i+j] %= 10
		}
	}

	// Remove leading zeros and reverse back
	return buildString(result)
}

// Helper function to reverse a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Helper function to build final string from result array
func buildString(result []int) string {
	// Find last non-zero index
	lastNonZero := len(result) - 1
	for lastNonZero >= 0 && result[lastNonZero] == 0 {
		lastNonZero--
	}

	// Build string from result array in correct order
	var sb strings.Builder
	for i := lastNonZero; i >= 0; i-- {
		sb.WriteRune(rune(result[i] + '0'))
	}

	return sb.String()
}
