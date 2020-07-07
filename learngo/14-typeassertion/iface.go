package main

import ."fmt"

func main() {
	typeassert()
}

func typeassert() {
	Println("-- type assertion")
	var i interface{} = "hello"
	
	s := i.(string); Println(s)
	s, ok := i.(string); Println(s, ok)
	f, ok := i.(float64); Println(f, ok)
	f = i.(float64) //panic
	Println(f)
}
