package main

import (
	"fmt"
	"time"
)

func Sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
	fmt.Println("Я очнулся")
}

func main() {
	var a int
	fmt.Print("Введите кол секунд: ")
	fmt.Scan(&a)
	Sleep(a)
}
