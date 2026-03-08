# Sorting Algorithm Benchmarks

This script will run benchmarks for the following sorting algorithms:

- Bubble Sort
- Selection Sort
- Insertion Sort
- Merge Sort
- Quick Sort

The results will be stored in the 'results' directory as CSV files, which can be used for further analysis and visualization.

## Hardware Specifications

- CPU: `AMD EPYC 7763 64-Core Processor`
- Cores: `2`
- RAM: `7.8Gi`
- OS: `Linux codespaces-363dba 6.8.0-1044-azure #50~22.04.1-Ubuntu SMP Wed Dec  3 15:13:22 UTC 2025 x86_64 GNU/Linux`

## Software Specifications

- Go Version: `go1.25.6`
- Sorting Algorithm Implementations:
  - Bubble Sort: [bubble.go](../modules/sort/bubble.go)
  - Selection Sort: [selection.go](../modules/sort/selection.go)
  - Insertion Sort: [insertion.go](../modules/sort/insertion.go)
  - Merge Sort: [merge.go](../modules/sort/merge.go)
  - Quick Sort: [quick.go](../modules/sort/quick.go)
- Benchmarking Code: [benchmark.go](../modules/benchmark.go)

## Benchmarking Methodology

- Each sorting algorithm will be tested with the same set of randomly generated arrays of varying sizes (e.g., 100, 1,000, 10,000, 100,000 elements).
- Each benchmark will be run multiple times (e.g., 5 iterations) to ensure consistency and to calculate average execution times.
- The results will be recorded in CSV format for easy analysis and visualization. Each CSV file will contain the following columns:
  - Algorithm: The name of the sorting algorithm being benchmarked.
  - Input Type: The type of input array (e.g., sorted_ascending, sorted_descending, random_shuffled).
  - Array Size: The size of the input array.
  - Best Time (ns): The best execution time in nanoseconds for the given array size across all iterations.
  - Worst Time (ns): The worst execution time in nanoseconds for the given array size across all iterations.
  - Average Time (ns): The average execution time in nanoseconds for the given array size across all iterations.
  - Median Time (ns): The median execution time in nanoseconds for the given array size across all iterations.
