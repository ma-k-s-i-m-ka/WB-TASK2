package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x, y}
}

func (p Point) Distance() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main() {
	var a, b float64
	fmt.Print("Введите X и Y: ")
	fmt.Scan(&a, &b)
	point := NewPoint(a, b)
	distance := point.Distance()
	fmt.Printf("Расстояние между точками X и Y (гипотенуза): %.1f\n", distance)
}
