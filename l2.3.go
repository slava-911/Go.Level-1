package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&num)
	if num < 100 || num > 999 {
		fmt.Println("Введено не трехзначное число")
		return
	}
	fmt.Printf("Количество единиц: %d\n", num%10)
	fmt.Printf("Количество десятков: %d\n", int(num/10)%10)
	fmt.Printf("Количество сотен: %d\n", int(num/100))
}
