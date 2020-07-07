package main

import ."fmt"

func main() {
	emptyreceiver()
}

func emptyreceiver() {
	Println("-- empty interface")
	var i interface{}; describe(i)
	i = 42; describe(i)
	i = "hi"; describe(i)
}

func describe(i interface{}){ Printf("(%v, %T)\n", i,i) }
