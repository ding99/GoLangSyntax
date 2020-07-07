package main

import (
	"fmt"
	"math"
)

func main() {
  parameters()
  closures()
  fibonc()
}

func parameters() {
  fmt.Println("-- function parameters")
  
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func closures() {
  fmt.Println("-- function closures")
  
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonc(){
  fmt.Println("-- Fibonacci closure")
  
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(i),",")
	}
	fmt.Println()
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
  return func(n int) int {
    x,y := 0,1
	for i := 0; i < n; i++{
	  x,y = y,x+y
	}
	return x
  }  
}
