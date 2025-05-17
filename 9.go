package main

import "fmt"

func main() {

	// IF else
	age := 4
	if age > 18 {
		fmt.Println(("You can vote"))
	} else if age > 10 {
		fmt.Println(("Bacchu hai tu"))
	} else {
		fmt.Println(("Sorry Bhaijan"))
	}

	// You can also assign variable in 'IF' block
	if aaa := 45; aaa == 45 {
		fmt.Println("Yes")
	}
}

// Commmand to run this file -
// go run 9.go
