package main

import "fmt"

func add(z int, y int) int {
	return z + y
}

func getLanguage() (string, string, string) {
	return "Go", "Javascript", "Python"
}

// func processIt(fn func(a int) int) {
// 	fn(1)
// }

func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	result := add(4, 6)
	fmt.Println(result)

	lan1, lan2, lan3 := getLanguage()
	fmt.Println(lan1, lan2, lan3)

	// fn := func(a int) int {
	// 	return 2
	// }

	fn := processIt()
	fn(6)
}
