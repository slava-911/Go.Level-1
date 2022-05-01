package main

import "fmt"

func main() {

	var n int
	numbers := map[int]int{
		0: 0,
		1: 1,
	}

	fmt.Print("Введите номер числа Фибоначчи: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Число введено неверно")
		return
	}
	fmt.Print("Число Фибоначчи: ", fibonacci2(n, numbers))
}

func fibonacci2(n int, m map[int]int) int {
	num, found := m[n]
	if found {
		return num
	} else {
		m[n] = fibonacci2(n-1, m) + fibonacci2(n-2, m)
		return m[n]
	}
}
