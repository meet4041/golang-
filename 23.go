package main

import (
	"fmt"
)

// Generic Slice
// func printSlice[T int | string](items []T) {
// func printSlice[T any](items []T) {
// func printSlice[T interface{}](items []T) {
func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// Integer Slice
// func printIntSlice(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// String Slice
// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func main() {

	// Generics

	nums := []int{1, 2, 3, 4, 5}
	// printIntSlice(nums)
	names := []string{"Meet", "Jainee"}
	// printStringSlice(names)

	fmt.Println(nums)
	fmt.Println(names)
}

// Commmand to run this file -
// go run 23.go
