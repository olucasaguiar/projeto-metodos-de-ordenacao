package sort

import "cmp"

func SelectionSort[T cmp.Ordered](array []T, size int) []T {
	for i := 0; i < size-1; i++ {
		minIdx := i
		for j := i + 1; j < size; j++ {
			if (array)[j] < (array)[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			(array)[i], (array)[minIdx] = (array)[minIdx], (array)[i]
		}
	}
	return array
}
