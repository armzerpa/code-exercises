package exercises

import (
	"sort"
	"strconv"
	"strings"
)

func SortPriorityTimestamp(arr []string) []string {
	sortedArr := make([]string, len(arr))
	copy(sortedArr, arr)

	sort.Slice(sortedArr, func(i, j int) bool {
		partsI := strings.Split(sortedArr[i], ":")
		partsJ := strings.Split(sortedArr[j], ":")

		priorityI, errI := strconv.Atoi(partsI[0])
		priorityJ, errJ := strconv.Atoi(partsJ[0])

		if errI != nil || errJ != nil {
			return false
		}

		if priorityI != priorityJ {
			return priorityI < priorityJ
		}

		timestampI, errI := strconv.Atoi(partsI[1])
		timestampJ, errJ := strconv.Atoi(partsJ[1])

		if errI != nil || errJ != nil {
			return false
		}

		return timestampI < timestampJ
	})

	return sortedArr
}
