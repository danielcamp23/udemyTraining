package main

import (
	"fmt"
	"math"
)

type square struct {
	val float64
}

func (mysquare square) area() float64 {
	return mysquare.val * mysquare.val
}

type circus struct {
	val float64
}

func (mycircus circus) area() float64 {
	return math.Pi * mycircus.val * mycircus.val
}

type shape interface {
	area() float64
}

func info(i shape) {
	fmt.Println(i)
	fmt.Println(i.area())
}

func main() {
	s := square{10}
	c := circus{8}
	info(s)
	info(c)
}

/*
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
*/
