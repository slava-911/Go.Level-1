package main

import "fmt"

func main() {
	var l, w float32
	fmt.Print("Введите длину: ")
	fmt.Scanln(&l)
	fmt.Print("Введите ширину: ")
	fmt.Scanln(&w)
	fmt.Printf("Площадь прямоугольника: %f\n", l*w)
}
