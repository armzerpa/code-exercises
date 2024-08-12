package exercises

func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLength := 0, 1

	for i := 0; i < len(s); i++ {
		// Check for odd-length palindromes
		len1 := expandAroundCenter(s, i, i)
		// Check for even-length palindromes
		len2 := expandAroundCenter(s, i, i+1)

		length := max(len1, len2)
		if length > maxLength {
			start = i - (length-1)/2
			maxLength = length
		}
	}

	return s[start : start+maxLength]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
