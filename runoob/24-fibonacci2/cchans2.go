package main

import ( ."fmt"; "time" )

func main(){
  c := make(chan int, 7)
  Println("capacity:", cap(c))
  
  go fibonacci(cap(c), c)
  
  for i := range c { Println("main: ", time.Now(), i) }
}

func fibonacci(n int, c chan int){
  x, y := 0, 1
  
  for i := 0; i < n; i++{
    c <- x
	Println("sub: ", time.Now())
	time.Sleep(50)
	x, y = y, x+y
  }
  
  close(c)  //program will be blocked if not close
}
