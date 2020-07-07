package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func main() {
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
  m := make(map[string]int)
  
  words := strings.Split(s," ")
  for _,v := range words { m[v]++ }
  
  return m
}
