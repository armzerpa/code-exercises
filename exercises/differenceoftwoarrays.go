package exercises

// Input: nums1 = [1,2,3], nums2 = [2,4,6]
// Output: [[1,3],[4,6]]

// Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
// Output: [[3],[]]
func findDifference(nums1 []int, nums2 []int) [][]int {

	var res1 []int
	for i1 := range nums1 {
		v1 := nums1[i1]
		if !valueIsInArray(nums2, v1) {
			res1 = append(res1, v1)
		}
	}

	var res2 []int
	for i2 := range nums2 {
		v2 := nums2[i2]
		if !valueIsInArray(nums1, v2) {
			res2 = append(res2, v2)
		}
	}
	res := [][]int{
		res1,
		res2,
	}
	return res
	//return [][]int{{1, 2}, {2}}
}

func valueIsInArray(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func findDifference2(nums1 []int, nums2 []int) [][]int {
	// Create maps to store unique elements from each array
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	// Add all elements from nums1 to set1
	for _, num := range nums1 {
		set1[num] = true
	}

	// Add all elements from nums2 to set2
	for _, num := range nums2 {
		set2[num] = true
	}

	// Create result arrays
	result := make([][]int, 2)
	result[0] = []int{} // Elements in nums1 but not in nums2
	result[1] = []int{} // Elements in nums2 but not in nums1

	// Find elements in nums1 but not in nums2
	for num := range set1 {
		if !set2[num] {
			result[0] = append(result[0], num)
		}
	}

	// Find elements in nums2 but not in nums1
	for num := range set2 {
		if !set1[num] {
			result[1] = append(result[1], num)
		}
	}

	return result
}
