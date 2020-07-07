package main

import "fmt"

func main() {
  capacity()
  nilslice()
  makeslice()
  appendslice()
}

func capacity(){
  fmt.Println("-- slice capacity")
  
  s := []int{2, 3, 5, 7, 11, 13}; printSlice(s)

  // Slice the slice to give it zero length.
  s = s[:0]; printSlice(s)
  // Extend its length.
  s = s[:4]; printSlice(s)
  // Drop its first one or two values.
  s = s[1:4]; printSlice(s)
  s = s[2:]; printSlice(s)
}

func nilslice() {
  fmt.Println("-- nil slice")
  
  var s []int; printSlice(s)
  if s == nil { fmt.Println("nil!") }
}

func makeslice(){
  fmt.Println("-- make slice")
  
  a := make([]int, 5); printSlice(a)
  b := make([]int, 0, 5); printSlice(b)
  c := b[:2]; printSlice(c)
  d := c[2:5]; printSlice(d)
}

func appendslice(){
  fmt.Println("-- append slice")
  
  var s []int; printSlice(s)
  s = append(s,0); printSlice(s)
  s = append(s,1); printSlice(s)
  s = append(s,2,3,4); printSlice(s)
  s = append(s,5,6); printSlice(s)
}

func printSlice(s []int) {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
