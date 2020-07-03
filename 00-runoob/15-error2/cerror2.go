//to check: Error is not used?

package main

import . "fmt"

type DIV_ERR struct { etype int; v1 int; v2 int } //v1-dividee, v2-divider

func (div_err DIV_ERR) Error() string{
  if 0 == div_err.etype { return "zero_divider"
  } else { return "unknown" }
}

func div(a int, b int) (int, *DIV_ERR){
  if b == 0 { return 0,&DIV_ERR{ 0,a,b }
  } else { return a/b, nil }
}

func main(){
  v,r:= div(100,2)
  if nil != r { Println("(1)fail:", r)
  } else { Println("(1)succeed:",v) }
  
  v,r = div(100,0)
  if nil != r { Println("(2)fail:", r)
  } else { Println("(2)succeed:",v) }
}
