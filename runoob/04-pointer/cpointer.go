package main

import . "fmt"

const max = 3

func main(){
  var ptr *int
  Printf("ptr value : %x\n", ptr)
  
  number := [max]int{5,6,7}
  var ptrs [max]*int  //pointer array
  
  for i := range number{
    ptrs[i] = &number[i]
  }
  
  for i,x := range ptrs{
    Printf("pointer array: index %d value %d address %d\n",i,*x,x)
  }
}