package main

import "fmt"

// By value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In func:", num)
// }

// By Referene
func changeNum(num *int) {
	*num = 5
	fmt.Println("In func:", *num)
}

func main() {

	// 'Pointers'
	num := 1

	// fmt.Println("Num: ", num)
	// fmt.Println("Memory address:", &num)

	changeNum(&num)
	fmt.Println("In Main:", num)
}

// Commmand to run this file -
// go run 18.go
