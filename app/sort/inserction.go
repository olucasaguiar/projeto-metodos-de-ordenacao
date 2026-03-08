package sort

import "cmp"

func InsertionSort[T cmp.Ordered](array []T, size int) []T {
	for i := 1; i < size; i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
	return array
}
