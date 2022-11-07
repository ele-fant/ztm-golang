//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
	"math"
)

// * Create a rectangle structure containing its coordinates
type Rectangle struct {
	A Point
	B Point
}

type Point struct {
	X float64
	Y float64
}

func width(r Rectangle) float64 {
	return math.Abs(r.B.X - r.A.X)
}

func length(r Rectangle) float64 {
	return math.Abs(r.B.Y - r.A.Y)
}

func lineLength(a Point, b Point) float64 {
	l := math.Sqrt(math.Pow((b.Y-a.Y), 2) + math.Pow((b.X-a.X), 2))
	return l
}

func area(r Rectangle) float64 {
	//p1 := Point{
	//	X: r.A.X,
	//	Y: r.B.Y,
	//}
	//a, b := lineLength(r.A, p1), lineLength(r.B, p1)
	return width(r) * length(r)
}

func perimiter(r Rectangle) float64 {
	//p1 := Point{
	//	X: r.A.X,
	//	Y: r.B.Y,
	//}
	//a, b := lineLength(r.A, p1), lineLength(r.B, p1)
	return width(r) + length(r)
}

func main() {
	p1, p2 := Point{X: 1, Y: 1}, Point{X: 5, Y: 4}
	rect := Rectangle{A: p1, B: p2}
	fmt.Println(rect)

	a := area(rect)
	fmt.Println("Area is:", a)

	perim := perimiter(rect)
	fmt.Println("Perimiter is:", perim)

	rect.B.X *= 2
	rect.A.X *= 2

	fmt.Println("With doubled parameter:")
	fmt.Println(rect)
	a = area(rect)
	fmt.Println("Area is:", a)

	perim = perimiter(rect)
	fmt.Println("Perimiter is:", perim)
}
