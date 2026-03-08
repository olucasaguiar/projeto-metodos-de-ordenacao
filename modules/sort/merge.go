package sort

import "cmp"

func MergeSort[T cmp.Ordered](array []T, size int) []T {
	if size <= 1 {
		return array
	}
	mergeSortHelper(&array, 0, size-1)
	return array
}

func mergeSortHelper[T cmp.Ordered](array *[]T, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergeSortHelper(array, left, mid)
		mergeSortHelper(array, mid+1, right)
		merge(array, left, mid, right)
	}
}

func merge[T cmp.Ordered](array *[]T, left, mid, right int) {
	leftSize := mid - left + 1
	rightSize := right - mid

	leftArr := make([]T, leftSize)
	rightArr := make([]T, rightSize)

	for i := range leftSize {
		leftArr[i] = (*array)[left+i]
	}
	for i := range rightSize {
		rightArr[i] = (*array)[mid+1+i]
	}

	i, j, k := 0, 0, left

	for i < leftSize && j < rightSize {
		if leftArr[i] <= rightArr[j] {
			(*array)[k] = leftArr[i]
			i++
		} else {
			(*array)[k] = rightArr[j]
			j++
		}
		k++
	}

	for i < leftSize {
		(*array)[k] = leftArr[i]
		i++
		k++
	}

	for j < rightSize {
		(*array)[k] = rightArr[j]
		j++
		k++
	}
}
