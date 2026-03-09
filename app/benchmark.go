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
	name    string
	sample  string
	size    int
	best    time.Duration
	worst   time.Duration
	average time.Duration
	median  time.Duration
}

func benchmark(fn func([]int, int) []int, array []int, size int) time.Duration {
	start := time.Now()
	fn(array, size)
	duration := time.Since(start)
	return duration
}

func calculateStatistics(durations []time.Duration) (best time.Duration, worst time.Duration, average time.Duration, median time.Duration) {
	if len(durations) == 0 {
		return 0, 0, 0, 0
	}

	best = durations[0]
	worst = durations[0]
	var total time.Duration

	for _, duration := range durations {
		if duration < best {
			best = duration
		}
		if duration > worst {
			worst = duration
		}
		total += duration
	}

	average = total / time.Duration(len(durations))

	sortedDurations := sort.QuickSort(durations, len(durations))

	if len(sortedDurations)%2 == 0 {
		mid1 := sortedDurations[len(sortedDurations)/2-1]
		mid2 := sortedDurations[len(sortedDurations)/2]
		median = (mid1 + mid2) / 2
	} else {
		median = sortedDurations[len(sortedDurations)/2]
	}

	return best, worst, average, median
}

func runSingleScenario(name string, description string, sortAlgorithm func([]int, int) []int, array []int, size int) TestResult {
	var duration []time.Duration

	for range EXECUTION_COUNT {
		sampleCopy := make([]int, len(array))
		copy(sampleCopy, array)
		duration = append(duration, benchmark(sortAlgorithm, sampleCopy, size))
	}

	best, worst, average, median := calculateStatistics(duration)

	return TestResult{
		name:    name,
		sample:  description,
		size:    size,
		best:    best,
		worst:   worst,
		average: average,
		median:  median,
	}
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
			results = append(results, result)
			fmt.Printf(" Concluído em %s\n", result.average)
		}
	}
	return results
}

func exportResults(results []TestResult) error {
	header := []string{"Algorithm", "Sample", "Size", "Best (ns)", "Worst (ns)", "Average (ns)", "Median (ns)"}
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
			fmt.Sprintf("%d", result.best.Nanoseconds()),
			fmt.Sprintf("%d", result.worst.Nanoseconds()),
			fmt.Sprintf("%d", result.average.Nanoseconds()),
			fmt.Sprintf("%d", result.median.Nanoseconds()),
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
