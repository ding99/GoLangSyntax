package main

import ( . "fmt"; "math" )

func main(){
  Println(math.Exp2(10)) //1024
  Print("a", "b", 1, 2, 3, "c", "d", "\n")
  Println("a","b",1,2,3,"c","d")
  
  var a = 1.5
  var c = 2
  Println(a,c)
  
  var i int
  var f float64
  var b bool
  var s string
  f1 := "Runoob"
  Printf("%v %v %v %q %q\n",  i, f, b, s, f1)
  
  var ( v1 int; v2 float32 )
  Printf("%v %v\n", v1, v2)
  Println("hello guy", f1)
  
  _,num,strs := numbers()
  Println(num,strs)
  
  constp()
}

func numbers()(int,int,string){
  a,b,c := 1,2,"strrrr"
  return a,b,c
}

func constp(){
  
  Println("-- const")
  const LENGTH, WIDTH = 10,5
  var area int
  const a,b,c = 1,false,"str1"
  
  area = LENGTH * WIDTH
  Printf("area : %d", area)
  Println()
  println(a,b,c)
}