package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var numbers []int64
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите набор целых чисел, каждое число на новой строке. Если введена не цифра, то ввод завершен.")
	for scanner.Scan() {
		num, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			break
		}
		numbers = append(numbers, num)
	}

	fmt.Println("Веденный набор чисел:")
	fmt.Println(numbers)
	sortInsert(numbers)
	fmt.Println("После сортировки:")
	for _, num := range numbers {
		fmt.Print(num, " ")
	}
}

func sortInsert(arr []int64) {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}
