package main

import (
	"fmt"
)

func main() {
	var str string

	fmt.Scan(&str)
	run := []rune(str)
	for i, j := 0, len(run)-1; i < j; i, j = i+1, j-1 {
		run[i], run[j] = run[j], run[i]
	}
	fmt.Println(string(run))
}
