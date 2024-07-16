package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func newPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := newPoint(3, 4)
	p2 := newPoint(6, 8)

	distance := p1.distance(p2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
