package main

import "fmt"

const age = 21

func main() {
	// const name = "Golang"
	// fmt.Println(name)
	fmt.Println(age)

	const (
		port = 5001
		host = "localhost"
	)
	fmt.Println(port, host)
}
