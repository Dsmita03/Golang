package main

import "fmt"

// for -> only construct in go for looping
func main() {

	//traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	//infinite loop
	k := 0
	for {
		if k >= 5 {
			break
		}
		fmt.Println(k)
		k++
	}

	//Range loop
	for l := range 10 {
		fmt.Println(l)
	}
}
