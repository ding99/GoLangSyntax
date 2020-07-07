package main

import ."fmt"

type I interface { M() }
type T struct { S string }

func (t *T) M() {
	if t == nil { Println("<nil>");
	} else { Println(t.S) }
}

func main() {
	nilreceiver()
}

func nilreceiver() {
	Println("-- interface with nill underlying values")
	var i I
	
	var t *T; i = t; describe(i); i.M()
	i = &T{"hey"}; describe(i); i.M()
}

func describe(i I){ Printf("(%v, %T)\n", i,i) }
