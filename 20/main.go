package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	words := strings.Fields(str)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	res := strings.Join(words, " ")

	fmt.Printf("Строка с перевернутыми словами: %s\n", res)
}
