package utilities

import "fmt"

// OddNumbers print some odd numbers
func OddNumbers() {
	fmt.Println("-- utilities : OddNumbers")
	ReverseRunes("srebmub olleH")
	for i := 0; i < 20; i++ {
		if i%2 == 1 {
			fmt.Print(i, ",")
		}
	}
	fmt.Println()
}
