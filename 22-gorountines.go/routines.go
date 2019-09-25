package main

import "fmt"
import "time"

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Second)
	}
}

func main(){

	f("direct")

	go f("goroutine")

	go func(msg string){
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
