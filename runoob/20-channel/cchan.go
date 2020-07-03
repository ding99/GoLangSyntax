package main

import . "fmt"

func main(){
  s := []int{ 7,2,8,-9,4,0,-1,10,11 }
  
  c := make(chan int, 100)
  go sum(s[:len(s)/2], c)
  go sum(s[len(s)/2:], c)
  x,y := <- c, <- c
  
  Println(x,y,x + y)
}

func sum(s []int, c chan int){
  sum := 0
  for _,v := range s { sum += v }
  c <- sum
}
