package exercises

func Exist(board [][]byte, word string) bool {
	// Check if the board is empty or word is empty
	if len(board) == 0 || len(word) == 0 {
		return false
	}

	// Iterate through each cell in the board
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			// If we find a starting match, attempt to find the full word
			if dfs2(board, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}

// Depth-first search to find the word
func dfs2(board [][]byte, i, j int, word string, index int) bool {
	// If we've matched all characters, we've found the word
	if index == len(word) {
		return true
	}

	// Check bounds and character mismatch
	if i < 0 || i >= len(board) ||
		j < 0 || j >= len(board[0]) ||
		board[i][j] != word[index] {
		return false
	}

	// Temporarily mark the current cell as visited
	temp := board[i][j]
	board[i][j] = '#'

	// Check all four directions
	found := dfs2(board, i+1, j, word, index+1) ||
		dfs2(board, i-1, j, word, index+1) ||
		dfs2(board, i, j+1, word, index+1) ||
		dfs2(board, i, j-1, word, index+1)

	// Restore the original character
	board[i][j] = temp

	return found
}
