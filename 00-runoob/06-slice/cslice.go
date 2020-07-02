package main

import . "fmt"

func main(){
  var numbers = make([]int,3,5)
  printSlice(numbers)
  
  var numbers1 []int
  printSlice(numbers1)
  
  if(numbers1 == nil){
    Printf("the slice is empty\n")
  }
  
  partial()
  appends()
}

func printSlice(x []int){
  Printf("len=%d cap=%d slice=%v\n", len(x),cap(x),x)
}

func partial(){
  Println("-- Partial Slice")
  
  numbers := []int{0,1,2,3,4,5,6,7,8}
  printSlice(numbers)
  
  Println("numbers ==", numbers)
  
  Println("numbers[1:4] ==", numbers[1:4])
  Println("numbers[:3] ==", numbers[:3])
  Println("numbers[4:] ==", numbers[4:])
  
  numbers1 := make([]int,0,5)
  printSlice(numbers1)
  
  numbers2 := numbers[:2]
  printSlice(numbers2)
  
  numbers3 := numbers[2:5]
  printSlice(numbers3)
}

func appends(){
  Println("-- append and copy")
  
  var numbers []int
  printSlice(numbers)
  
  Println("append 1")
  numbers = append(numbers,1)
  printSlice(numbers)
  
  Println("append 2-4")
  numbers = append(numbers,2,3,4)
  printSlice(numbers)
  
  numbers1 := make([]int,len(numbers),(cap(numbers))*2)
  copy(numbers1,numbers)
  printSlice(numbers1)
  
  Println("append-append")
  var arr1 = []int{1,2,3}
  var arr2 = []int{4,5,6}
  var arr3 = []int{7,8,9}
  var s = append(append(arr1,arr2...),arr3...)
  printSlice(s)
}