package exercises

func generateParenthesis(n int) []string {
	result := []string{}
	backtrack2(&result, "", 0, 0, n)
	return result
}

func backtrack2(result *[]string, current string, open int, close int, max int) {
	// Base case: if the current string length is 2 * max, we've created a valid combination
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	// Add an open parenthesis if we haven't used all open parentheses
	if open < max {
		backtrack2(result, current+"(", open+1, close, max)
	}

	// Add a close parenthesis if it's valid (fewer close than open)
	if close < open {
		backtrack2(result, current+")", open, close+1, max)
	}
}
