package main

import (
	."fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func main() {
	methodnonstruct()
	methodabs(); functionabs()
	methodpointer(); functionpointer()
}

func methodnonstruct() {
	Println("-- method on non-struct")
	f := MyFloat(-math.Sqrt2); Println(f.Abs())
}

func (f MyFloat) Abs() float64 {
	if f < 0 { return float64(-f) }
	return float64(f)
}

func methodabs(){
	Println("-- method abs")
	v := Vertex{3, 4}; Println(v.Abs())
}

func functionabs() {
	Println("-- function abs")
	v := Vertex{3, 4}; Println(Abs(v))
}

func (v Vertex) Abs() float64 { return math.Sqrt(v.X*v.X + v.Y*v.Y) }
func Abs(v Vertex) float64 { return math.Sqrt(v.X*v.X + v.Y*v.Y) }

func methodpointer() {
	Println("-- method with pointer reveiver")
	v := Vertex{3, 4}; v.Scale(10); Println(v.Abs())
}

func functionpointer() {
	Println("-- function with pointer parameter")
	v := Vertex{3, 4}; Scale(&v, 10); Println(Abs(v))
}

func (v *Vertex) Scale(f float64){ v.X *= f; v.Y *= f }
func Scale(v *Vertex, f float64) { v.X *= f; v.Y *= f }
