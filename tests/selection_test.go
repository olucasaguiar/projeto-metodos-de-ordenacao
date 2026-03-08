package sort

import (
	"testing"

	"github.com/olucasaguiar/projeto-metodos-de-ordenacao/app/sort"
)

func TestSelectionSort(t *testing.T) {
	type args struct {
		array []int
		size  int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "ordered ascendence",
			args: args{
				array: []int{1, 2, 3, 4, 5},
				size:  5,
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "ordered descendence",
			args: args{
				array: []int{5, 4, 3, 2, 1},
				size:  5,
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "shuffle",
			args: args{
				array: []int{3, 1, 4, 1, 5, 9, 2, 6},
				size:  8,
			},
			expected: []int{1, 1, 2, 3, 4, 5, 6, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.SelectionSort(tt.args.array, tt.args.size)
			for i, v := range tt.args.array {
				if v != tt.expected[i] {
					t.Errorf("SelectionSort() = %v, expected %v", tt.args.array, tt.expected)
					return
				}
			}
		})
	}
}
