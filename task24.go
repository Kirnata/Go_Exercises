package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

func GetDistance(p1, p2 Point) float64 {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	delta := math.Sqrt(float64(deltaX*deltaX + deltaY*deltaY))
	return delta
}

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point
// с инкапсулированными параметрами x,y и конструктором
func main() {
	var x1, x2, y1, y2 int

	fmt.Print("Coordinates of the first point: ")
	_, err := fmt.Scan(&x1, &y1)
	if err != nil {
		return
	}
	fmt.Print("Coordinates of the second point: ")
	_, err = fmt.Scan(&x2, &y2)
	if err != nil {
		return
	}

	point1 := NewPoint(x1, y1)
	point2 := NewPoint(x2, y2)
	d := GetDistance(point1, point2)
	fmt.Printf("Distance: %f", d)
}
