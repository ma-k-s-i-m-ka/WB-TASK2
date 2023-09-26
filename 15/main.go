package main

import "fmt"

var justString string

func createHugeString(v int) string {
	s := "hellohellohellohellohello"
	for i := 0; i < v; i++ {
		s += s
	}
	return s
}
func someFunc() {
	v := createHugeString(10)
	justString = v[:100]
	fmt.Println(justString)
	v = ""
}

func main() {
	someFunc()
}
