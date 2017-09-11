package main

import "math"
import "fmt"

type abser interface {
	Abs() float64
}

type myfloat float64

type vertex struct {
	x, y float64
}

func (f myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.x*v.x)
}

func main() {
	var i abser
	f := myfloat(-math.Sqrt2)
	v := vertex{3, 4}

	fmt.Println("1) v", v)
	fmt.Println("2) abs from v", v.Abs())

	fmt.Println("3) abs from f", f.Abs())

	i = f
	fmt.Println("4) abs from i", i.Abs())

	i = v
	fmt.Println("5) abs from i", i.Abs())

	i = &v
	fmt.Println("6) abs from i", i.Abs())

}

/*
package main

import "math"
import "fmt"

type abser interface {
	Abs() float64
}

type myfloat float64

type vertex struct {
	x, y float64
}

func (f myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.x*v.x)
}

func main() {
	var i abser
	f := myfloat(-math.Sqrt2)
	v := vertex{3, 4}

	fmt.Println("1) v", v)
	fmt.Println("2) abs from v", v.Abs())

	fmt.Println("3) abs from f", f.Abs())

	i = f
	fmt.Println("4) abs from i", i.Abs())

	i = &v
	fmt.Println("6) abs from i", i.Abs())

	i = v//these two lines do not run beacause i is a Vertex (not *Vertex)
	fmt.Println("5) abs from i", i.Abs())//and does NOT implement Abser.

}
*/
