package main

import ( ."fmt"; ."time" )

func main(){
	selects()
	defaultselect()
}
	
func selects() {
	Println("-- select")
	c := make(chan int)
	quit := make(chan int)
	
	go func() {
		for i := 0; i < 10; i++ { Print(<-c, ",") }
		Println()
		quit <- 0
	}()
	
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: x, y = y, x+y
		case <-quit: Println("quit"); return
		}
	}
}

func defaultselect() {
	Println("-- Default selectoin")
	
	tick := Tick(100 * Millisecond)
	boom := After(500 * Millisecond)
	for {
		select {
		case <-tick:
			Println("tick.")
		case <-boom:
			Println("BOOM!")
			return
		default:
			Println("    .")
			Sleep(50 * Millisecond)
		}
	}
}
