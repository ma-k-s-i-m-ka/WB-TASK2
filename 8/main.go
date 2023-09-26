package main

import (
	"fmt"
	"strconv"
)

func SetBit(num int64, pos uint) int64 {
	return num | (1 << pos)
}

func ClearBit(num int64, pos uint) int64 {
	return num &^ (1 << pos)
}

func main() {
	var num, index int64
	fmt.Print("Введите целое число: ")
	fmt.Scan(&num)
	binaryStr := strconv.FormatInt(num, 2)
	fmt.Printf("Число %d в бинарном виде: %s\n", num, binaryStr)
	fmt.Print("Введите индекс бита: ")
	fmt.Scan(&index)
	i := uint(index)

	num = SetBit(num, i)
	binaryStr = strconv.FormatInt(num, 2)
	fmt.Printf("Число после установки %d-го бита в 1: %d, в бинарном виде: %s\n ", i, num, binaryStr)

	num = ClearBit(num, i)
	binaryStr = strconv.FormatInt(num, 2)
	fmt.Printf("Число после установки %d-го бита в 0: %d, в бинарном виде: %s\n ", i, num, binaryStr)
}
