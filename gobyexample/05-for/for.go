package main

import "fmt"

func main(){
	i := 1
	for i <= 3{
		fmt.Println(i)
		i = i + 1
	}
	
	for j:= 7; j <= 9; j++{
		fmt.Println(j)
	}

	for{
		fmt.Println("loop 1")
		break;
	}

	k := 0;
	for{
		fmt.Println("loop 2")
		k++;
		if(k > 3){
			break
		}
	}
	fmt.Println("k = ", k)
	
	for n := 0; n < 6; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}