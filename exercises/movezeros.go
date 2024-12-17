package exercises

func movezeros(nums []int) {
	// Initialize the position where non-zero elements should be placed
	nonZeroIndex := 0

	// First pass: Move all non-zero elements to the front
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex] = nums[i]
			nonZeroIndex++
		}
	}

	// Second pass: Fill the remaining elements with zeros
	for i := nonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}
