package main

import . "fmt"

func main(){
  const (
    a = iota  //0
	b
	c
	d = "ha" //independant, while iota++
	e
	f = 100 // iota++
	g       //100, iota++
	h = iota  //7, recovery
	i)
	Println(a,b,c,d,e,f,g,h,i)
	
	ex1()
}

const(
  i = 1 << iota
  j = 3 << iota
  k
  l
)

func ex1(){
  Println("i=", i)
  Println("j=", j)
  Println("k=", k)
  Println("l=", l)
}