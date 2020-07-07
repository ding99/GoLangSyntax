package main

import ."fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func getSqrt(x float64, i float64) (float64,float64){
  remain := (i * i - x) / (2 * i)
  i -= remain
  
  if remain > 0 { return getSqrt(x, i)
  } else { return i, remain }
}

func Sqrt(x float64) (float64, error) {
  if(x < 0){
  	return x, ErrNegativeSqrt(x)
  }
  i,_ := getSqrt(x,x)
  return i, nil
}

func main() {
	Println(Sqrt(2))
	Println(Sqrt(-2))
}
