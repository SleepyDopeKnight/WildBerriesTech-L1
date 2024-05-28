package main

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64 // инкапсуляция происходит благодря регистру
}

func main() {
	point1 := newPoint(3, 4)
	point2 := newPoint(9, 6)
	fmt.Println(RangeBetweenPoints(point1, point2))
}

func newPoint(a, b float64) *Point { // конструктор
	return &Point{x: a, y: b}
}

func RangeBetweenPoints(point1 *Point, point2 *Point) float64 {
	return math.Sqrt(math.Pow(point2.x-point1.x, 2) + math.Pow(point2.y-point1.y, 2))
} // формула вычисления AB = √(xb - xa)2 + (yb - ya)2
