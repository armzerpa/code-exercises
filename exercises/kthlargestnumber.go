package exercises

import (
	"math/rand"
)

func reArrengeArrayPartition(nums []int, left, right int) int {
	pivot := nums[right]
	pIndex := left
	for i := left; i < right; i++ {
		if nums[i] <= pivot {
			nums[i], nums[pIndex] = nums[pIndex], nums[i]
			pIndex++
		}
	}
	nums[pIndex], nums[right] = nums[right], nums[pIndex]
	//fmt.Println("nums4: ", nums)
	return pIndex
}

func quickSelection(nums []int, left, right, k int) int {
	//fmt.Println("left: ", left, "right: ", right, "k: ", k)
	//fmt.Println("nums: ", nums)
	if left == right {
		return nums[left]
	}

	// Choose a random pivot index
	pivotIndex := left + rand.Intn(right-left+1)
	//fmt.Println("pivotIndex: ", pivotIndex)
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	//fmt.Println("nums2: ", nums)

	// Partition the array
	pIndex := reArrengeArrayPartition(nums, left, right)
	//fmt.Println("pIndex3: ", pIndex)
	//fmt.Println("nums3: ", nums)

	// Determine whether the k-th largest is on the left or right
	if k == pIndex {
		return nums[k]
	} else if k < pIndex {
		return quickSelection(nums, left, pIndex-1, k)
	} else {
		return quickSelection(nums, pIndex+1, right, k)
	}
}

func FindKthLargest(nums []int, k int) int {
	return quickSelection(nums, 0, len(nums)-1, len(nums)-k)
}
