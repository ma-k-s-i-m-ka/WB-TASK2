package main

import "fmt"

type Human struct {
	Gender string
	Name   string
	Age    int
}

type Action struct {
	Human
}

func (h Human) Example() {
	fmt.Println(h.Gender, h.Name, h.Age)
}

func main() {
	human := Human{
		Gender: "Male",
		Name:   "Maks",
		Age:    21,
	}
	action := Action{Human: human}
	action.Example()
}
