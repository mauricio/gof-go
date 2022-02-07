package gof_go

import "testing"

type Sorter func(a []int)

func MergeSort(a []int) {
	// implementation
}

func QuickSort(a []int) {
	// implementation
}

func TestSorter(t *testing.T) {
	var sorter Sorter
	sorter = MergeSort
	sorter([]int{10, 7, 5, 2, 4})
}
