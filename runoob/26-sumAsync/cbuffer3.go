package main
import ( ."fmt"; ."time" )

func main(){
  s := []int{ 7,2,8,-9,4,0 }
  c := make(chan int, 4)
  
  Println("Go 1: [0,3]"); go sum(s[:len(s)/2], c)
  Println("Go 1: [3,6]"); go sum(s[len(s)/2:], c)
  Println("Go 2: [0,3]"); go sum(s[:len(s)/2], c)
  Println("Go 2: [3,6]"); go sum(s[len(s)/2:], c)
  
  Println("Go 3: start waiting..."); Sleep(1000 * Millisecond)
  Println("Go 3: waited 1000 ms")

  aa := <- c; bb := <- c; Println(aa); Println(bb)
  x,y := <- c, <- c; Println(x,y,x + y)
}

func sum(s []int, c chan int){
  sum := 0
  for _,v := range s { sum += v }
  Printf("SUB: sum: %#v\n", sum)
  c <- sum
  Println("SUB: after channel pro")
}
