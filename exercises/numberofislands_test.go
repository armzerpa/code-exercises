package exercises

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name:     "Empty Grid",
			grid:     [][]byte{},
			expected: 0,
		},
		{
			name:     "Single Row Empty Grid",
			grid:     [][]byte{{}},
			expected: 0,
		},
		{
			name: "Single Island",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "Multiple Islands",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			name: "No Islands",
			grid: [][]byte{
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 0,
		},
		{
			name: "Single Cell Island",
			grid: [][]byte{
				{'1'},
			},
			expected: 1,
		},
		{
			name: "Diagonal Islands",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '0', '0'},
				{'1', '0', '1'},
			},
			expected: 4,
		},
		{
			name: "Complex Pattern",
			grid: [][]byte{
				{'1', '1', '0', '1', '0'},
				{'1', '0', '0', '0', '0'},
				{'0', '0', '1', '0', '1'},
				{'1', '0', '0', '1', '1'},
			},
			expected: 5,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the grid to preserve original for error reporting
			gridCopy := makeGridCopy(tt.grid)

			result := numIslands(tt.grid)

			if result != tt.expected {
				t.Errorf("Test case %s failed: expected %d islands, got %d\nOriginal grid:\n%s",
					tt.name, tt.expected, result, formatGrid(gridCopy))
			}
		})
	}
}

// Helper function to create a deep copy of the grid
func makeGridCopy(grid [][]byte) [][]byte {
	if len(grid) == 0 {
		return [][]byte{}
	}

	copy := make([][]byte, len(grid))
	for i := range grid {
		copy[i] = make([]byte, len(grid[i]))
		for j := range grid[i] {
			copy[i][j] = grid[i][j]
		}
	}
	return copy
}

// Helper function to format grid for error messages
func formatGrid(grid [][]byte) string {
	if len(grid) == 0 {
		return "[]"
	}

	result := "\n"
	for _, row := range grid {
		result += "["
		for _, cell := range row {
			result += string(cell) + " "
		}
		result = result[:len(result)-1] + "]\n"
	}
	return result
}

// Benchmark function to measure performance
func BenchmarkNumIslands(b *testing.B) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gridCopy := makeGridCopy(grid)
		numIslands(gridCopy)
	}
}
