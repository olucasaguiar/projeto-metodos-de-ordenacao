package main

import (
	"fmt"
	"os"
	"time"

	"github.com/olucasaguiar/projeto-metodos-de-ordenacao/modules/sort"
	"github.com/olucasaguiar/projeto-metodos-de-ordenacao/modules/utils"
)

type TestScenario struct {
	name         string
	arrayFactory func(int) []int
	sizes        []int
}

func benchmark(fn func([]int, int) []int, array []int, size int) time.Duration {
	start := time.Now()
	fn(array, size)
	duration := time.Since(start)
	return duration
}

func runSingleScenario(name string, sortAlgorithm func([]int, int) []int, array []int, size int) {
	duration := benchmark(sortAlgorithm, array, size)
	println(name, "took", duration.String())
}

func runAllScenarios(name string, sortAlgorithm func([]int, int) []int) {
	variations := []int{1_000, 10_000, 100_000}
	scenarios := []TestScenario{
		{"Sorted Array (asc)", utils.BuildAscendingSortedArray, variations},
		{"Sorted Array (desc)", utils.BuildDescendingSortedArray, variations},
		{"Random Array (shuffled)", utils.BuildRandomArray, variations},
	}
	for _, scenario := range scenarios {
		for _, size := range scenario.sizes {
			runSingleScenario(fmt.Sprintf("Running %s with %s of size %d...\n", name, scenario.name, size), sortAlgorithm, scenario.arrayFactory(size), size)
		}
	}
}

func main() {
	helpMessage := "Usage: go run app.go [test_case]\n" +
		"Where [test_case] can be one of the following:\n" +
		"  - bubble: Run benchmarks for Bubble Sort\n" +
		"  - selection: Run benchmarks for Selection Sort\n" +
		"  - insertion: Run benchmarks for Insertion Sort\n" +
		"  - merge: Run benchmarks for Merge Sort\n" +
		"  - quick: Run benchmarks for Quick Sort\n" +
		"Example: go run app.go bubble"

	if len(os.Args) < 2 {
		fmt.Println(helpMessage)
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments. Please provide only one test case: bubble, selection, insertion, merge, quick.")
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	testCase := os.Args[1]

	switch testCase {
	case "bubble":
		runAllScenarios("Bubble Sort", sort.BubbleSort)
	case "selection":
		runAllScenarios("Selection Sort", sort.SelectionSort)
	case "insertion":
		runAllScenarios("Insertion Sort", sort.InsertionSort)
	case "merge":
		runAllScenarios("Merge Sort", sort.MergeSort)
	case "quick":
		runAllScenarios("Quick Sort", sort.QuickSort)
	default:
		println("Invalid test case. Please use one of the following: bubble, selection, insertion, merge, quick.")
		return
	}

}
