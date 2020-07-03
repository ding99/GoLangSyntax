package main

import . "fmt"

func main(){
  Println("OUTER: start")
  
  defer func(){
    Println("OUTER: prepare recover")
	if err := recover(); err != nil { Println("OUTER: %#v-%#v\n", "outer", err)
	  //err was already caught by upper. no exception here, but run defer and its coming code
    } else { Println("OUTER: nothing to do") }
	Println("OUTER: finish recover")
  }()
  
  Println("OUTER: before exception")
  f()
  Println("OUTER: after exception")

  defer func() { Println("OUTER: defer after exception") }()
}

func f() {
  Println("INNER: start")
  defer func() { Println("INNER: defer before recover") }()
  
  defer func() {
    Println("INNER: prepare recover")
	if err := recover(); err != nil { Printf("INNER: %#v-%#v\n", "inner", err) } //err from panic
	Println("INNER: finish recover")
  }()
  
  defer func() { Println("INNER: defer after recover / before exception") }()
  panic("INNER: Exception Info")
  defer func() { Println("INNER: defer after exception") }()
  
  Println("INNER: statements after exception") //recover ,or not catch. not run from here
}

/* result:
OUTER: start
OUTER: before exception
INNER: start
INNER: defer after recover / before exception
INNER: prepare recover
INNER: "inner"-"INNER: Exception Info"
INNER: finish recover
INNER: defer before recover
OUTER: after exception
OUTER: defer after exception
OUTER: prepare recover
OUTER: nothing to do
OUTER: finish recover
*/