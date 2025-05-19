package main

import "fmt"

func sum(nums ...int) int {
	// func sum(nums ...interface{}) int {		// anytype acceptable
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}

func main() {

	// 'VARIADIC Fucnction'

	nums := []int{3, 4, 5, 6, 88}
	result := sum(nums...)
	fmt.Println(result)

}

// Commmand to run this file -
// go run 16.go
