package main

import "fmt"

// int->0 || string ="" || bool=false

func main() {

	// 'ARRAY'

	var nums [5]int
	fmt.Println(len(nums)) // Length of an array
	nums[2] = 3            // Assign value array
	fmt.Println(nums)      // Printing whole array

	// To declare it in single line
	example1 := [3]int{1, 2, 3}
	fmt.Println(example1)

	// 2D array
	example2 := [2][2]int{{1, 1}, {2, 2}}
	fmt.Println(example2)

	// Benifits:
	// - Fixed size predictable
	// - Memory optimization
	// - constant time access
}

// Commmand to run this file -
// go run 11.go
