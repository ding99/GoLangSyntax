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
  
  var a []string = strings.Split(s," ")
  for i := 0; i < len(a); i++{
    _, ok := m[a[i]]
	if ok {
	  m[a[i]]++
	} else { 
      m[a[i]] = 1
	}
  }
  return m
}
