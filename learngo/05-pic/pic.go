package main

import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
  p := make([][]uint8,dy)
  for i := 0; i < dy; i++ {
    p[i] = make([]uint8,dx)
	for j := 0; j < dx; j++ {
	  p[i][j] = uint8((i + j) / 2)
	}
  }
  
  fmt.Println("row length:",len(p), ", capacity:",cap(p))
  return p
}

func main() {
	pic.Show(Pic)
}
