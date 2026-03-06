package sort

func MergeSort(array []int, size int) {
	if size <= 1 {
		return
	}
	mergeSortHelper(array, 0, size-1)
}

func mergeSortHelper(array []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergeSortHelper(array, left, mid)
		mergeSortHelper(array, mid+1, right)
		merge(array, left, mid, right)
	}
}

func merge(array []int, left, mid, right int) {
	leftSize := mid - left + 1
	rightSize := right - mid

	leftArr := make([]int, leftSize)
	rightArr := make([]int, rightSize)

	for i := 0; i < leftSize; i++ {
		leftArr[i] = (array)[left+i]
	}
	for i := 0; i < rightSize; i++ {
		rightArr[i] = (array)[mid+1+i]
	}

	i, j, k := 0, 0, left

	for i < leftSize && j < rightSize {
		if leftArr[i] <= rightArr[j] {
			(array)[k] = leftArr[i]
			i++
		} else {
			(array)[k] = rightArr[j]
			j++
		}
		k++
	}

	for i < leftSize {
		(array)[k] = leftArr[i]
		i++
		k++
	}

	for j < rightSize {
		(array)[k] = rightArr[j]
		j++
		k++
	}
}
