package main

import (
	"fmt"
	"math"
)

func main() {
	var s, d float64
	fmt.Print("Введите площадь круга: ")
	fmt.Scanln(&s)
	d = 2 * math.Sqrt(s/math.Pi)
	fmt.Printf("Диаметр окружности: %f\n", d)
	fmt.Printf("Длина окружности: %f\n", math.Pi*d)
}
