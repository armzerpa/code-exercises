package exercises

import "sort"

func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var current []int

	// Sort the candidates to optimize the backtracking
	sort.Ints(candidates)

	backtrack(candidates, target, 0, current, &result)

	return result
}

func backtrack(candidates []int, target int, start int, current []int, result *[][]int) {
	if target == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}

		current = append(current, candidates[i])
		backtrack(candidates, target-candidates[i], i, current, result)
		current = current[:len(current)-1]
	}
}
