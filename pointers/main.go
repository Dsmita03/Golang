package main

import "fmt"

// by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("num in changeNum:", num)
// }

// by reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("num in changeNum:", *num)
}

func main() {
	num := 1
	// changeNum(num)
	changeNum(&num)
	fmt.Println("Memory address", &num)
	fmt.Println("num in main:", num)
}
