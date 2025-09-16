package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch
	i := 5
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	//switch with multiple conditions
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	//type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		case bool:
			fmt.Println("I'm a bool")
		default:
			fmt.Println("I'm something else")
		}
	}
	whoAmI(50) //whatever passed through interface according to that it will print the type
}
