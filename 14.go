package main

import "fmt"

// iterating over data strucure
func main() {
	// nums := []int{1, 2, 4}

	// Simple For Loop
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// Range
	// for i, num := range nums {
	// 	fmt.Println("Index :", i, ", Value: ", num)
	// }

	// m := map[string]string{"fname": "Meet", "lname": "Gandhi"}

	// for k := range m { 	// only key
	// for k, v := range m { // key and value
	// 	fmt.Println(k, v)
	// }

	// unicode
	// starting byte of rune
	for k, v := range "golang" { // key and value
		fmt.Println(k, v)
	}
}

// Commmand to run this file -
// go run 14.go
