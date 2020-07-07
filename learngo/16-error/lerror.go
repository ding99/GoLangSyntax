package main

import ( ."fmt"; ."time" )

type MyError struct {
	When Time
	What string
}

func (e *MyError) Error() string {
	return Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{ Now(),"it didn't work" }
}

func main() {
	if err := run(); err != nil {
		Println(err)
	}
}
