package main
import . "fmt"

func main(){
  var a int = 4
  var b int32
  var c float32
  var ptr *int
  
  Printf("line 1 - a variable type = %T\n", a)
  Printf("line 2 - b variable type = %T\n", b)
  Printf("line 3 - c variable type = %T\n", c)

  ptr = &a
  Printf("value of a %d\n", a)
  Printf("*prt %d\n", *ptr)
}