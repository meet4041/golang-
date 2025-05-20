package main

import "fmt"

func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}

func main() {

	// 'Clousers'

	ans := counter()
	fmt.Println(ans())
	fmt.Println(ans())

}

// Commmand to run this file -
// go run 17.go
