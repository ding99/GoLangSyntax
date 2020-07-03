package main

import . "fmt"

type DivideError struct { dividee int; divider int }

func (de *DivideError) Error() string{
  strFormat :=`
  Cannot proceed, the divider is zero.
  dividee: %d
  dovider: 0
`
  return Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errMsg string){
  if varDivider == 0 {
    dData := DivideError{
	  dividee: varDividee,
	  divider: varDivider,
	}
	errorMsg := dData.Error()
	return 0, errorMsg
  } else {
    return varDividee / varDivider, ""
  }
}

func main(){
  if result, errorMsg := Divide(100,10); errorMsg == "" {
    Println("100/10 = ", result)
  }
  if _,errorMsg := Divide(100,0); errorMsg != "" {
    Println("errorMsg is: ", errorMsg)
  }
}
