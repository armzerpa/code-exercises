package exercises

func MissingNumber(nums []int) int {
	n := len(nums)
	result := n

	for i := 0; i < n; i++ {
		result ^= i ^ nums[i]
	}

	return result
}
