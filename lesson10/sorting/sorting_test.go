package sorting

import "testing"

var slc = []int{999, 33, 43453, 2, 99334, 0, 33, 666, 7777, 55555, 999999, 111111}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(slc)
	}
}

// func BenchmarkQuicksort(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Quicksort(slc)
// 	}
// }

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(slc)
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort2(slc)
	}
}

func BenchmarkBubbleSort3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort3(slc)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectSort(slc)
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShellSort(slc)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort(slc)
	}
}
