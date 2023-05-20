package sorting

import (
	"YadroTestWork/internal/utils/structures"
	"sort"
)

func Sort(array []string, compares ...func(a, b structures.Pair[string, int]) bool) (result structures.Result) {
	compare := func(a, b structures.Pair[string, int]) bool {
		return a.First < b.First
	}
	if len(compares) != 0 {
		compare = compares[0]
	}

	if len(array) == 0 {
		return
	}
	sort.Strings(array)
	count := structures.Pair[string, int]{First: array[0], Second: 1}

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
	sort.SliceStable(result, func(i, j int) bool {
		return compare(result[i], result[j])
	})
	return
}
