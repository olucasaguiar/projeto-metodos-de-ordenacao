package sort

import (
	"cmp"
)

func BubbleSort[T cmp.Ordered](array []T, size int) []T {
	for i := 0; i < size-1; i++ {
		swapped := false
		for j := 0; j < size-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return array
}
