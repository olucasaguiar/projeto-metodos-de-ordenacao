package main

import (
	"time"

	"github.com/olucasaguiar/projeto-metodos-de-ordenacao/modules"
	"github.com/olucasaguiar/projeto-metodos-de-ordenacao/modules/sort"
)

func benchmark(fn func([]int, int), array []int, size int) time.Duration {
	start := time.Now()
	fn(array, size)
	duration := time.Since(start)
	return duration
}

func main() {
	tamanho := 10000
	vector := modules.VetorAleatorio(tamanho)

	durationBubbleSort := benchmark(sort.BubbleSort, vector, tamanho)
	durationInsertionSort := benchmark(sort.InsertionSort, vector, tamanho)
	durationSelectionSort := benchmark(sort.SelectionSort, vector, tamanho)
	durationMergeSort := benchmark(sort.MergeSort, vector, tamanho)
	durationQuickSort := benchmark(sort.QuickSort, vector, tamanho)

	println("Bubble Sort:", durationBubbleSort.Nanoseconds()/1e6, "ms")
	println("Insertion Sort:", durationInsertionSort.Nanoseconds()/1e6, "ms")
	println("Selection Sort:", durationSelectionSort.Nanoseconds()/1e6, "ms")
	println("Merge Sort:", durationMergeSort.Nanoseconds()/1e6, "ms")
	println("Quick Sort:", durationQuickSort.Nanoseconds()/1e6, "ms")
}
