package main

import ( ."fmt"; ."math" )

type I interface { M() }
type T struct { S string }
type F float64

func (t *T) M() { Println(t.S) }
func (f F) M() { Println(f) }

func main() {
	implicit()
}

func implicit() {
	Println("-- implemented implicitly")
	var i I

	i = &T{"Hello"}; describe(i); i.M()
	i = F(Pi); describe(i); i.M()
}

func describe(i I){ Printf("(%v, %T)\n", i,i) }
