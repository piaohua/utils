package utils

import (
	"math/rand"
	"testing"
	"time"
)

func TestShellSort(t *testing.T) {
	var arr []int
	arr = []int{9, 92, 8, 5, 21, 1, 5, 3, 2, 13}
	ShellSort(arr)
	t.Log("ShellSort:", arr)
	arr = []int{9, 92, 8, 5, 21, 1, 5, 3, 2, 13}
	ShellSort(arr)
	t.Log("ShellSort:", arr)
}

func TestSort(t *testing.T) {
	arr := []int{9, 92, 8, 5, 21, 1, 5, 3, 2, 12}
	InsertionSort(arr)
	t.Log("InsertionSort:", arr)

	arr = []int{9, 92, 8, 5, 21, 1, 5, 3, 2, 16}
	InsertSort(arr)
	t.Log("InsertSort:", arr)

	arr = []int{9, 92, 8, 5, 21, 1, 5, 3, 2, 13}
	ShellSort(arr)
	t.Log("ShellSort:", arr)

	arr = []int{9, 92, 8, 5, 21, 1, 5, 3, 2, 10}
	BubbleSort(arr)
	t.Log("BubbleSort:", arr)

	arr = []int{9, 92, 8, 5, 21, 1, 5, 3, 2, 11}
	QuickSort(arr)
	t.Log("QuickSort:", arr)

	arr = []int{9, 92, 8, 5, 21, 1, 5, 3, 2, 22}
	SelectionSort(arr)
	t.Log("SelectionSort:", arr)

	arr = []int{9, 92, 8, 5, 21, 1, 5, 3, 2, 93}
	arr = MergeSort(arr)
	t.Log("MergeSort:", arr)

}

func TestRadixSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	sz := make([]int, 30)
	for i := 0; i < 30; i++ {
		sz[i] = rand.Intn(90) + 10
	}
	t.Log(sz)
	RadixSort(sz)
	t.Log(sz)
}
