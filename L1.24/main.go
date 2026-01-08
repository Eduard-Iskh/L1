package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (point *Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow(point.x-other.x, 2) + math.Pow(point.y-other.y, 2))
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}
func main() {
	point1 := NewPoint(-1, 3)
	point2 := NewPoint(6, 2)
	fmt.Println(point1.Distance(*point2))
}
