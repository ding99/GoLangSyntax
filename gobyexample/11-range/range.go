package main

import "fmt"

func main(){

	nums := []int{1,2,3}
	fmt.Println("nums: ", nums)
	
	sum := 0
	for _,num := range nums{
		sum += num
	}
	fmt.Println("sum:", sum)
	
	for i, num := range nums{
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	fmt.Println("-- 1st map")
	kvs := map[string]string{"a":"apple","b":"banana"}
	for k,v := range kvs{
		fmt.Printf("%s -> %s\n", k, v)
	}
	
		for k := range kvs {
		fmt.Println("key:", k)
	}

	fmt.Println("-- 2nd map")
	kvt := map[string]int{"a":7,"b":12}
	for k,v := range kvt{
		fmt.Printf("%s -> %d\n", k, v)
	}
	
	for k := range kvt {
		fmt.Println("key:", k)
	}
	
	fmt.Println("-- 3rd map")
	kvu := map[int]string{3:"a",7:"b"}
	for k,v := range kvu{
		fmt.Printf("%d -> %s\n", k, v)
	}
	
	for k := range kvu {
		fmt.Println("key:", k)
	}
	
	fmt.Println("--")
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	for i, c := range "study" {
		fmt.Printf("%d -> %Xh\n", i, c)
	}
}