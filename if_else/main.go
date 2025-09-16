package main

import "fmt"

func main() {
	//if-else
	age := 20
	if age >= 18 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}

	//logical operations
	var role = "admin"
	var hasPermission = true

	if role == "admin" && hasPermission {
		fmt.Println("You can access the admin panel.")
	}

	//we can declare a variable inside  if constructor
	if age := 20; age >= 18 {
		fmt.Println("You are adult.", age)
	} else if age >= 12 {
		fmt.Println("Teenager", age)
	}

	// Golang does not have ternary operator
	// condition ? value1 : value2

}
