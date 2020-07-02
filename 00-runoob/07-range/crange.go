package main

import . "fmt"

func main(){
  sum_normal(10)
  sum_range()
}

func sum_normal(max int){
  Println("== sum, normal")
  Print("-- method # 1: ")
  sum := 0
  for i := 0; i <= max; i++ { sum += i }
  Printf("sum %d, count %d\n", sum, max)

  Print("-- method # 2: ")
  sum = 1
  for ; sum <= max; { sum += sum }
  Printf("sum %d, upper %d\n", sum, max)
  
  Print("-- method # 3: ")
  sum = 1
  for sum <= max { sum += sum }
  Printf("sum %d, upper %d\n", sum, max)
}

func sum_range(){
  Println("== sum, range")

  Print("-- range # 1: ")
  strings := []string{"google","facebook"}
  for i, s := range strings{ Printf("%d %s, ", i,s) }
  Println()
  
  Print("-- range # 2: ")
  numbers := [6]int{1,2,3,5}
  for i, x := range numbers{ Printf("%d-%d, ", i, x) }
  Println()
  
  Print("-- range # 3: ")
  // to get the sum of a slice using range. Similar to array
  nums := []int{2,3,4}
  sum := 0
  for _, num := range nums{ sum += num }
  Println(sum)
  
  Print("-- range-map: ")
  kvs := map[string]string{ "a":"apple", "b":"banana" }
  for k, v := range kvs{ Printf("%s -> %s, ",k,v) }
  Println()
  
  Print("-- range-unicode: ")
  for i, c := range "google" { Printf("%d-%d, ",i,c) }
  Println()
  
  Print("-- simple loop: ")
  nums = []int{1,2,3,4}
  length := 0
  for range nums { length++ }
  Println(length)
  
  Print("-- pair loop: ")
  for i,num := range nums { Printf("%d-%d, ", i,num) }
  Println()
}