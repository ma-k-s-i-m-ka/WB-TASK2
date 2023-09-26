package main

import "fmt"

func main() {
	m := make(map[string]int)
	for _, word := range []string{"cat", "cat", "dog", "cat", "tree"} {
		m[word]++
	}
	for k, v := range m {
		fmt.Printf("\n Key: %s, Value: %d", k, v)
	}
}
