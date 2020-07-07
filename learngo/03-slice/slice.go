package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
  slices()
  sliceliterals()
  bounds()
}

func slices(){
  fmt.Println("-- slices")

  names := [4]string{ "John","Paul","George","Ringo" }
  fmt.Println(names)
  
  a := names[0:2]
  b := names[1:3]
  fmt.Println(a,b)

  b[0] = "XXX"
  fmt.Println(a,b)
  fmt.Println(names)
}

func sliceliterals(){
  fmt.Println("-- slices-literals")

  q := []int{2, 3, 5, 7, 11, 13}
  fmt.Println(q)

  r := []bool{true, false, true, true, false, true}
  fmt.Println(r)

  s := []struct { i int; b bool
  }{ {2, true},{3, false},{5, true},{7, true},{11, false},{13, true} }

  fmt.Println(s)
}

func bounds(){
  fmt.Println("-- slice bounds")
  
  s := []int{2,3,5,7,11,13 }
  s = s[1:4];fmt.Println(s);
  s = s[:2];fmt.Println(s);
  s = s[1:];fmt.Println(s);
}