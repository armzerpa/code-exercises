package exercises

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// Iterate through each cell in the grid
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				// Found an island, increment count and sink it
				count++
				dfs(grid, i, j)
			}
		}
	}

	return count
}

// DFS to mark all connected land cells as visited
func dfs(grid [][]byte, row, col int) {
	// Check boundary conditions and if current cell is land
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) || grid[row][col] != '1' {
		return
	}

	// Mark current cell as visited by changing it to '0'
	grid[row][col] = '0'

	// Recursively check all adjacent cells
	dfs(grid, row+1, col) // down
	dfs(grid, row-1, col) // up
	dfs(grid, row, col+1) // right
	dfs(grid, row, col-1) // left
}
