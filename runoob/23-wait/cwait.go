package main

import ( ."fmt"; ."time" )

func main(){
  ch := make(chan int)
  go say2("sub", ch)
  say("main")
  Println(<- ch) //waiting ch from sub(say2)
}

func say(s string){
  for i := 0; i < 5; i++{
    Sleep(50 * Millisecond)
	Println(s, (i + 1) * 50)
  }
}

func say2(s string, ch chan int){
  for i := 0; i < 5; i++{
    Sleep(100 * Millisecond)
	Println(s, (i + 1) * 100)
  }
  ch <- 0
  close(ch)
}
