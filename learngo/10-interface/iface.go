package main

import (
	."fmt"
)

func main() {
	implicit()
}

func implicit() {
	Println("-- implemented implicitly")
	
	var i I = T{"Hello"}
	i.M()
}

type I interface { M() }
type T struct { S string }

func (t T) M() { Println(t.S) }
