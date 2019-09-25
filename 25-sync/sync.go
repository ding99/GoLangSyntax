package main

import "fmt"
import "time"

func worker(done chan bool){
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done 1")
	done <- true
	time.Sleep(time.Second)
	fmt.Println("...done 2")
	done <- false
	done <- true
	time.Sleep(time.Second)
	fmt.Println("...done 3")
	done <- false
	done <- true
}

func main(){

	done := make(chan bool, 2)
	go worker(done)

	fmt.Println(<-done)
	fmt.Println(<-done)
	fmt.Println(<-done)
	fmt.Println(<-done)
	fmt.Println(<-done)
}


