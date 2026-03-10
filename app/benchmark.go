package main

import (
	"fmt"
	"os"
	"time"

	"github.com/olucasaguiar/projeto-metodos-de-ordenacao/app/sort"
	"github.com/olucasaguiar/projeto-metodos-de-ordenacao/app/utils"
)

const (
	REPORT_DIRNAME      = "results"
	REPORT_FILENAME     = "benchmarks.csv"
	EXECUTION_COUNT     = 5
	BUBBLE_SORT         = "bubble"
	SELECTION_SORT      = "selection"
	INSERTION_SORT      = "insertion"
	MERGE_SORT          = "merge"
	QUICK_SORT          = "quick"
	ALL_SORT_ALGORITHMS = BUBBLE_SORT + "," + SELECTION_SORT + "," + INSERTION_SORT + "," + MERGE_SORT + "," + QUICK_SORT
	SORTED_ASCENDING    = "sorted_ascending"
	SORTED_DESCENDING   = "sorted_descending"
	RANDOM_SHUFFLED     = "random_shuffled"
)

type TestScenario struct {
	name         string
	arrayFactory func(int) []int
	sizes        []int
}

type TestResult struct {
	name      string
	sample    string
	size      int
	iteration int
	duration  time.Duration
}

func benchmark(fn func([]int, int) []int, array []int, size int) time.Duration {
	start := time.Now()
	fn(array, size)
	duration := time.Since(start)
	return duration
}

func runSingleScenario(name string, description string, sortAlgorithm func([]int, int) []int, array []int, size int) []TestResult {
	var results []TestResult

	for i := range EXECUTION_COUNT {
		sampleCopy := make([]int, len(array))
		copy(sampleCopy, array)
		duration := benchmark(sortAlgorithm, sampleCopy, size)
		result := TestResult{
			name:      name,
			sample:    description,
			size:      size,
			iteration: i,
			duration:  duration,
		}
		results = append(results, result)
	}

	return results
}

func runAllScenarios(algorithmName string, sortAlgorithm func([]int, int) []int) []TestResult {
	sampleSizes := []int{1_000, 10_000, 100_000}
	scenarios := []TestScenario{
		{SORTED_ASCENDING, utils.BuildAscendingSortedArray, sampleSizes},
		{SORTED_DESCENDING, utils.BuildDescendingSortedArray, sampleSizes},
		{RANDOM_SHUFFLED, utils.BuildRandomArray, sampleSizes},
	}
	results := make([]TestResult, 0)
	for _, scenario := range scenarios {
		for _, size := range scenario.sizes {
			fmt.Printf("Executando %s para %s de tamanho %d...", algorithmName, scenario.name, size)
			result := runSingleScenario(algorithmName, scenario.name, sortAlgorithm, scenario.arrayFactory(size), size)
			results = append(results, result...)
			fmt.Println(" Concluído.")
		}
	}
	return results
}

func exportResults(results []TestResult) error {
	header := []string{"Algorithm", "Sample", "Size", "Iteration", "Duration (ns)"}
	report, err := utils.NewReport(REPORT_DIRNAME, REPORT_FILENAME, header, false)
	if err != nil {
		return err
	}

	var csvLines [][]string
	for _, result := range results {
		csvLines = append(csvLines, []string{
			result.name,
			result.sample,
			fmt.Sprintf("%d", result.size),
			fmt.Sprintf("%d", result.iteration),
			fmt.Sprintf("%d", result.duration.Nanoseconds()),
		})
	}
	return report.WriteLines(csvLines...)
}

func main() {
	helpMessage := "Uso: go run app/benchmark.go [caso_teste]\n" +
		"Onde [caso_teste] pode ser um dos seguintes:\n" +
		"  - " + BUBBLE_SORT + ": Executar benchmarks para Bubble Sort\n" +
		"  - " + SELECTION_SORT + ": Executar benchmarks para Selection Sort\n" +
		"  - " + INSERTION_SORT + ": Executar benchmarks para Insertion Sort\n" +
		"  - " + MERGE_SORT + ": Executar benchmarks para Merge Sort\n" +
		"  - " + QUICK_SORT + ": Executar benchmarks para Quick Sort\n" +
		"Exemplo: go run app/benchmark.go " + BUBBLE_SORT

	if len(os.Args) < 2 {
		fmt.Println(helpMessage)
		os.Exit(0)
	}
	if len(os.Args) > 2 {
		fmt.Println("Entrada inválida. Por favor, forneça apenas um caso de teste: " + ALL_SORT_ALGORITHMS + ".")
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	testCase := os.Args[1]

	fmt.Println("Iniciando benchmarks...")

	var results []TestResult
	switch testCase {
	case BUBBLE_SORT:
		results = runAllScenarios(BUBBLE_SORT, sort.BubbleSort)
	case SELECTION_SORT:
		results = runAllScenarios(SELECTION_SORT, sort.SelectionSort)
	case INSERTION_SORT:
		results = runAllScenarios(INSERTION_SORT, sort.InsertionSort)
	case MERGE_SORT:
		results = runAllScenarios(MERGE_SORT, sort.MergeSort)
	case QUICK_SORT:
		results = runAllScenarios(QUICK_SORT, sort.QuickSort)
	default:
		fmt.Println("Caso de teste inválido. Por favor, use um dos seguintes: " + ALL_SORT_ALGORITHMS + ".")
		os.Exit(1)
	}

	fmt.Println("Exportando resultados para benchmarks.csv...")
	err := exportResults(results)
	if err != nil {
		fmt.Println("Erro ao exportar resultados:", err)
		os.Exit(1)
	}
	fmt.Println("Benchmarking concluído com sucesso.")
}
