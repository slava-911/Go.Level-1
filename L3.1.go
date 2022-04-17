package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, res float32
	var op string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		ProcExit("Число введено неверно")
	}
	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, !): ")
	fmt.Scanln(&op)
	if op == "" {
		ProcExit("Операция не введена")
	} else if op == "!" {
		if a < 0 {
			ProcExit("Для операции факториал число должно быть положительным")
		}
		res = float32(Factorial(uint(a)))
	} else {
		fmt.Print("Введите второе число: ")
		_, err := fmt.Scanln(&b)
		if err != nil {
			ProcExit("Число введено неверно")
		}
		switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			if b == 0 {
				ProcExit("На ноль делить нельзя")
			}
			res = a / b
		case "^":
			res = Power(a, int(b))
		default:
			ProcExit("Операция введена неверно")
		}
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}

func Factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Power(x float32, n int) float32 {
	var m int
	var y float32
	if n < 0 {
		m = -n
	} else {
		m = n
	}
	y = 1
	for i := 1; i <= m; i++ {
		y *= x
	}
	if n < 0 {
		y = 1 / y
	}
	return y
}

func ProcExit(text string) {
	fmt.Println(text)
	os.Exit(1)
}
