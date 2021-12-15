package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}
type Line struct {
	a, b Point
}

func (l Line) Distance() float64 {
	return math.Hypot(l.a.x-l.b.x, l.a.y-l.b.y)
}

type Path []Point

func (p Path) Distance() (sum float64) {
	for i := 1; i < len(p); i++ {
		sum += Line{p[i], p[i-1]}.Distance()
	}
	return
}

func main() {
	side := Line{Point{1, 2}, Point{4, 6}}
	fmt.Println(side.Distance())
	journey := Path{Point{1, 2}, Point{4, 6}, Point{7, 10}, Point{10, 14}}
	fmt.Println(journey.Distance())
}
