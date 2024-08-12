package exercises

import (
	"fmt"
	"math"
)

func CoinChange(coins []int, amount int) int {
	fmt.Println("coins:", coins, "amount:", amount)
	// Create a slice to store the minimum number of coins for each amount
	dp := make([]int, amount+1)

	// Initialize all values to a number larger than possible coin count
	for i := range dp {
		dp[i] = amount + 1
	}

	// Base case: 0 coins needed to make amount 0
	dp[0] = 0

	// Build up the solution for all amounts from 1 to target amount
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
				//fmt.Println("dp 1:", dp)
			}
		}
	}
	//fmt.Println("dp result:", dp)

	// If dp[amount] is still amount+1, it means no solution was found
	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
