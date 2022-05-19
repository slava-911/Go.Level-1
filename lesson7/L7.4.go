package main

import (
	"fmt"
	"io"
)

type ByteSlice []byte

func (bs *ByteSlice) Write(arg []byte) (n int, err error) {
	*bs = append(*bs, arg...)
	return len(arg), nil
}

func (bs *ByteSlice) Read(arg []byte) (n int, err error) {
	n = copy(arg, *bs)
	*bs = (*bs)[n:]
	if n == 0 {
		err = io.EOF
	}
	return
}

func main() {

	var bs ByteSlice
	str := "io.Reader, io.Writer"

	fmt.Println("Исходная строка:\n", str)
	io.WriteString(&bs, str)

	fmt.Println("Слайс байтов после записи io.WriteString:\n", bs)

	ReadSlice, err := io.ReadAll(&bs)
	if err != nil {
		fmt.Println("Ошибка!", err.Error())
	}
	fmt.Println("Строка из слайса байтов после считывания io.ReadAll:\n", string(ReadSlice))

}
