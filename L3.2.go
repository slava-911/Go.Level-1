package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Print("Введите целое число, до которого необходимо найти простые числа: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Число введено неверно")
		return
	}

	numbers := Eratosthenes(n)
	fmt.Println(numbers)
	fmt.Print("Number of prime numbers generated: ", len(numbers))
}

func Eratosthenes(n int) []int {
	arr := make([]bool, n+1)
	var numbers []int

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

	return numbers
}
