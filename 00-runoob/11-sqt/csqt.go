package main

import . "fmt"

func main(){
  Println(getSqrt(2))
//  Println(getSqrt(3))
}

func getSqrt(x float64) float64 {
  i,_ := sqrt(x,x)
  Println()
  return i  
}

func sqrt(x float64,i float64)(float64,float64){
  remain := (i*i-x)/(2*i)
  i = i - remain
  Printf(" (%g/%g)",remain,i)
  if remain > 0 { return sqrt(x,i)
  } else { return i, remain }
}
