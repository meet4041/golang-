package main

import "fmt"

func main() {

	// 'MAPS' -> hash, object, dict

	// Creating Map
	m := make(map[string]string)

	// Setting element in a map
	m["name"] = "Meet Gandhi"
	m["techstack"] = "ML"

	fmt.Println(m)
	fmt.Println(m["name"])
	fmt.Println(m["techstack"])

	// If key doesn't exist in map it returns 0 value
	fmt.Println(m["phone"])

	i := make(map[string]int)
	fmt.Println(i["num"])
	fmt.Println(len(i))

	// Delete element from stack
	delete(m, "techstack")
	fmt.Println(m)

	// Clear the map
	clear(i)
	fmt.Println(i)

	// z := make(map[string]int)
	z := map[string]int{"price": 40, "age": 22}
	fmt.Println(z)

	v, ok := z["age"]
	fmt.Println(v)
	if ok {
		fmt.Println("Ok")
	} else {
		fmt.Println("Not Ok")
	}
}

// Commmand to run this file -
// go run 13.go
