package primes

import (
	"errors"
	"math"
)

var ErrNotPositive = errors.New("input number must be positive")

func Eratosthenes(n int) (numbers []int, err error) {

	if n < 0 {
		return numbers, ErrNotPositive
	}

	arr := make([]bool, n+1)

	for i := 2; i <= int(math.Sqrt(float64(n)+1)); i++ {
		if arr[i] == false {
			for j := i * i; j <= n; j += i {
				arr[j] = true
			}
		}
	}

	for i, isComposite := range arr {
		if i > 1 && !isComposite {
			numbers = append(numbers, i)
		}
	}

	return numbers, nil
}
