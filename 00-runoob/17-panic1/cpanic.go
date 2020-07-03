package main

import . "fmt"

func main(){
  Println("MAIN: 1")
  defer func(){
    if msg := recover(); msg != nil { Println("Printing exception --> ", msg) }
  }()
  Println("MAIN: 2")
  pic()
  Println("MAIN: 3")
}

func pic() {
  Println("PIC: 1")
  panic("PIC: panic...")
  Println("PIC: 2")
}

/* result:
MAIN: 1
MAIN: 2
PIC: 1
Printing exception -->  PIC: panic...
*/
