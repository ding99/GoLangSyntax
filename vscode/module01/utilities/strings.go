package utilities

import "fmt"

//ReverseRunes reverse a string
func ReverseRunes(input string) string {
	fmt.Println("-- utilities : ReverseRunes")
	ret := []rune(input)
	for i, j := 0, len(ret)-1; i < len(ret)/2; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return string(ret)
}
