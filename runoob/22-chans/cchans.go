package main

import . "fmt"

func main(){
  c := make(chan int, 11)
  Println("capacity:", cap(c))
  
  go fibonacci(cap(c), c)
  
  for i := range c { Println(i) }
}

func fibonacci(n int, c chan int){
  x, y := 0, 1
  for i := 0; i < n; i++{ c <- x; x, y = y, x+y }
  close(c)  //program will be blocked if not close
}
