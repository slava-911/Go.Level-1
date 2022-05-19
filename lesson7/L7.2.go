package main

import (
	"errors"
	"fmt"
)

var EnterErr = errors.New("неверный ввод")
var PosNumErr = errors.New("число должно быть положительным")
var DivErr = errors.New("на ноль делить нельзя")

func main() {
	res, err := calc()
	if err != nil {
		fmt.Println("Ошибка!", err.Error())
		return
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}

func calc() (float64, error) {
	var a, b, res float64
	var op string

	res = 0
	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		return res, fmt.Errorf("Первое число: %w", EnterErr)
	}
	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, !): ")
	fmt.Scanln(&op)
	if op == "" {
		return res, errors.New("Операция не введена")
	} else if op == "!" {
		if a < 0 {
			return res, fmt.Errorf("Операция факториал: %w", PosNumErr)
		}
		res = float64(Factorial(uint(a)))
	} else {
		fmt.Print("Введите второе число: ")
		_, err := fmt.Scanln(&b)
		if err != nil {
			return res, fmt.Errorf("Второе число: %w", EnterErr)
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
				return res, fmt.Errorf("Операция деление: %w", DivErr)
			}
			res = a / b
		case "^":
			res = Power(a, int(b))
		default:
			return res, fmt.Errorf("Операция: %w", EnterErr)
		}
	}

	return res, nil
}

func Factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func Power(x float64, n int) float64 {
	var m int
	var y float64
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
