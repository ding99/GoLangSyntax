package main

import "fmt"
import "math"

const s string = "Constant"

func main(){
	fmt.Println(s)
	
	const n = 500000000
	
	const d = 3e20 / n
	fmt.Println(d)
	
	fmt.Println(int64(d))
	fmt.Println(int32(n))
	
	fmt.Println(math.Sin(n))
	fmt.Println(math.Cos(n))
}