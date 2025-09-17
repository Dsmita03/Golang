package main

import (
	"fmt"
	"maps"
)

// maps -> hash,dictionary,object
func main() {
	// creating map
	m := make(map[string]string) // map[keyType]valueType

	// setting an element
	m["name"] = "Golang"
	m["area"] = "backend"
	// get an element
	fmt.Println(m["name"], m["area"])
	fmt.Println(len(m)) //length

	// deleting an element
	delete(m, "area")
	fmt.Println(m)

	//we can also create and initialize a map in one line
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(m1)

	m2 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(maps.Equal(m1, m2))

	// IMP: if key does not exist, it returns zero value
}
