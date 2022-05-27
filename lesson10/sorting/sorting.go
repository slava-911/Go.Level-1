package sorting

func InsertionSort(ar []int) {
	n := len(ar)
	var i, j int
	for i = 1; i < n; i++ {
		temp := ar[i]
		for j = i; j > 0 && ar[j-1] > temp; j-- {
			ar[j] = ar[j-1]
		}
		ar[j] = temp
	}
}

func Quicksort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

func partition(ar []int) int {
	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		// swap(ar, left, right)
		ar[left], ar[right] = ar[right], ar[left]
	}
}

func BubbleSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if ar[i] > ar[j] {
				// swap(ar, i, j)
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
	}
}

func BubbleSort2(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := len(ar) - 1; j > i; j-- {
			if ar[j-1] > ar[j] {
				// swap(ar, j-1, j)
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
	}
}

func BubbleSort3(ar []int) {
	for i := 0; i < len(ar); i++ {
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				// swap(ar, j-1, j)
				ar[j-1], ar[j] = ar[j], ar[j-1]
			}
		}
	}
}

func SelectSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		min := i
		for j := i + 1; j < len(ar); j++ {
			if ar[min] > ar[j] {
				min = j
			}
		}

		if min != i {
			// swap(ar, i, min)
			ar[i], ar[min] = ar[min], ar[i]
		}
	}
}

func ShellSort(ar []int) {
	for gap := len(ar) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(ar); i++ {
			x := ar[i]
			j := i
			for ; j >= gap && ar[j-gap] > x; j -= gap {
				ar[j] = ar[j-gap]
			}
			ar[j] = x
		}
	}
}

func HeapSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	heapIt(ar)

	ar[0], ar[len(ar)-1] = ar[len(ar)-1], ar[0]

	HeapSort(ar[:len(ar)-1])
}

func heapIt(ar []int) {

	// left = 2*i + 1
	// right = 2*i + 2
	// i = 0
	// мы каждый раз рассматриваем 3 ноды - рут и 2 листа
	// и по кажлому листу повторяем рекурсивно heapIt(ar[1:]), heapIt(ar[2:])
	if len(ar) < 2 {
		return
	}

	if len(ar) == 2 {
		if ar[0] < ar[1] {
			ar[0], ar[1] = ar[1], ar[0]
		}
		return
	}

	if len(ar) > 3 {
		heapIt(ar[1:])
		heapIt(ar[2:])
	}

	if ar[1] > ar[2] {
		if ar[0] < ar[1] {
			ar[0], ar[1] = ar[1], ar[0]
		}
	} else {
		if ar[0] < ar[2] {
			ar[0], ar[2] = ar[2], ar[0]
		}
	}
}
