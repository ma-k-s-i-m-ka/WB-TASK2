package main

import (
	"fmt"
	"reflect"
)

func main() {

	var integer int = 42
	var str string = "Hello, world!"
	var booL bool = true
	var ch chan int = make(chan int)

	var myInterface interface{}

	myInterface = integer
	if res, ok := myInterface.(int); ok {
		fmt.Printf("Тип переменной через Ok: int: %d\n", res)
	}

	switchType(myInterface)
	myInterface = str
	switchType(myInterface)
	myInterface = booL
	switchType(myInterface)
	myInterface = ch
	switchType(myInterface)

	reflectType(integer, str, booL, ch)
}

func switchType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Тип переменной через switch: int, Значение: %d\n", v)
	case string:
		fmt.Printf("Тип переменной через switch: string, Значение: %s\n", v)
	case bool:
		fmt.Printf("Тип переменной через switch: bool, Значение: %t\n", v)
	case chan int:
		fmt.Printf("Тип переменной через switch: chan int, Значение: %v\n", v)
	default:
		fmt.Printf("Неизвестный тип\n")
	}
}
func reflectType(a ...interface{}) {
	for _, i := range a {
		t := reflect.TypeOf(i)
		fmt.Printf("Тип переменной через reflect: %s\n", t)
	}
}
