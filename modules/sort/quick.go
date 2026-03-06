package sort

func QuickSort(array []int, size int) []int {
	if size <= 1 {
		return array
	}
	quickSortHelper(&array, 0, size-1)
	return array
}

func quickSortHelper(array *[]int, low, high int) {
	if low < high {
		pi := partition(array, low, high)
		quickSortHelper(array, low, pi-1)
		quickSortHelper(array, pi+1, high)
	}
}

func partition(array *[]int, low, high int) int {
	pivot := (*array)[high]
	i := low - 1

	for j := low; j < high; j++ {
		if (*array)[j] < pivot {
			i++
			(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
		}
	}
	(*array)[i+1], (*array)[high] = (*array)[high], (*array)[i+1]
	return i + 1
}
