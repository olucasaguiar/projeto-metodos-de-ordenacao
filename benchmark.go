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

type TestResult struct {
	name        string
	description string
	size        int
	duration    time.Duration
}

func benchmark(fn func([]int, int) []int, array []int, size int) time.Duration {
	start := time.Now()
	fn(array, size)
	duration := time.Since(start)
	return duration
}

func runSingleScenario(name string, description string, sortAlgorithm func([]int, int) []int, array []int, size int) TestResult {
	duration := benchmark(sortAlgorithm, array, size)
	return TestResult{
		name:        name,
		description: description,
		size:        size,
		duration:    duration,
	}
}

func runAllScenarios(algorithmName string, sortAlgorithm func([]int, int) []int) []TestResult {
	sampleSizes := []int{1_000, 10_000, 100_000}
	scenarios := []TestScenario{
		{"Sorted Array (asc)", utils.BuildAscendingSortedArray, sampleSizes},
		{"Sorted Array (desc)", utils.BuildDescendingSortedArray, sampleSizes},
		{"Random Array (shuffled)", utils.BuildRandomArray, sampleSizes},
	}
	results := make([]TestResult, 0)
	for _, scenario := range scenarios {
		for _, size := range scenario.sizes {
			fmt.Printf("Running %s for %s of size %d...", algorithmName, scenario.name, size)
			sampleDescription := fmt.Sprintf("Execution using %s of size %d", scenario.name, size)
			result := runSingleScenario(algorithmName, sampleDescription, sortAlgorithm, scenario.arrayFactory(size), size)
			fmt.Printf(" Done in %s\n", result.duration)
			results = append(results, result)
		}
	}
	return results
}

func exportResults(filename string, results []TestResult) error {
	report, err := utils.NewReport("results", filename, true)
	if err != nil {
		return err
	}

	err = report.WriteHeader("Algorithm", "Description", "Size", "Duration (ns)")
	if err != nil {
		return err
	}

	var csvLines [][]string
	for _, result := range results {
		csvLines = append(csvLines, []string{
			result.name,
			result.description,
			fmt.Sprintf("%d", result.size),
			fmt.Sprintf("%d", result.duration.Nanoseconds()),
		})
	}
	return report.WriteLines(csvLines...)
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
		os.Exit(0)
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments. Please provide only one test case: bubble, selection, insertion, merge, quick.")
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	testCase := os.Args[1]

	fmt.Println("Starting benchmarks...")

	var results []TestResult
	switch testCase {
	case "bubble":
		results = runAllScenarios("Bubble Sort", sort.BubbleSort)
	case "selection":
		results = runAllScenarios("Selection Sort", sort.SelectionSort)
	case "insertion":
		results = runAllScenarios("Insertion Sort", sort.InsertionSort)
	case "merge":
		results = runAllScenarios("Merge Sort", sort.MergeSort)
	case "quick":
		results = runAllScenarios("Quick Sort", sort.QuickSort)
	default:
		fmt.Println("Invalid test case. Please use one of the following: bubble, selection, insertion, merge, quick.")
		os.Exit(1)
	}

	fmt.Printf("Exporting results to %s_benchmarks.csv...\n", testCase)
	err := exportResults(fmt.Sprintf("%s_benchmarks.csv", testCase), results)
	if err != nil {
		fmt.Println("Error exporting results:", err)
		os.Exit(1)
	}
	fmt.Println("Benchmarking completed successfully.")
}
