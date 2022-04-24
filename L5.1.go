package main

import "fmt"

func main() {

	var n int

	fmt.Print("Введите номер числа Фибоначчи: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Число введено неверно")
		return
	}
	fmt.Print("Число Фибоначчи: ", fibonacci(n))
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
