package sort

import (
	"YadroTestWork/internal/utils/structures"
	"sort"
)

func Sort(array []string) []structures.Pair[string, int] {
	if len(array) == 0 {
		return make([]structures.Pair[string, int], 0)
	}
	sort.Strings(array)
	count := structures.Pair[string, int]{First: array[0], Second: 1}

	result := make([]structures.Pair[string, int], 0)

	for i := 1; i < len(array); i++ {
		if count.First == array[i] {
			count.Second++
		} else {
			result = append(result, count)
			count.First = array[i]
			count.Second = 1
		}
	}
	result = append(result, count)

	return result
}
