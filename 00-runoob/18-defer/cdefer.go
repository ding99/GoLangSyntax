package main

import . "fmt"

func main(){
  Println("MAIN: 1")
  defer Println("MAIN: defer 1")
  Println("MAIN: 2")
  defer Println("MAIN: defer 2")
  defer Println("MAIN: defer 3")
  fool()
}

func fool() {
  defer Println("FOOL: defer 1")
  defer Println("FOOL: defer 2")
  Println("FOOL: 1")
  defer Println("FOOL: defer 3")
  Println("FOOL: 2")
}

/* Result should be:
MAIN: 1
MAIN: 2
FOOL: 1
FOOL: 2
FOOL: defer 3
FOOL: defer 2
FOOL: defer 1
MAIN: defer 3
MAIN: defer 2
MAIN: defer 1
*/