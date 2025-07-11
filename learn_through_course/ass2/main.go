package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	sq := square{sideLength: 10}
	tr := triangle{height: 10, base: 10}

	printArea(sq)
	printArea(tr)

}
