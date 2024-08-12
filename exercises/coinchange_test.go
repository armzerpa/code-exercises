package exercises

import "testing"

// Unit tests for coinChange function
func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},   // 11 = 5 + 5 + 1
		{[]int{2}, 3, -1},         // Cannot make amount 3 with coin 2
		{[]int{1}, 0, 0},          // Amount 0 requires 0 coins
		{[]int{1}, 2, 2},          // 2 = 1 + 1
		{[]int{1, 2, 5}, 0, 0},    // Amount 0 requires 0 coins
		{[]int{4, 5, 10}, 7, -1},  // Cannot make amount 7 with coins 2, 5, 10
		{[]int{3, 7, 405}, 17, 3}, // 17 = 7 + 7 + 3
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := CoinChange(tt.coins, tt.amount)
			if got != tt.want {
				t.Errorf("coinChange(%v, %d) = %d; want %d", tt.coins, tt.amount, got, tt.want)
			}
		})
	}
}
