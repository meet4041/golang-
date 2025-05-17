package main

import (
	"fmt"
)

func main() {

	// 'SWITCH'

	// Simple Switch
	// i := 2
	// switch i {
	// case 1:
	// 	fmt.Println(i)
	// case 2:
	// 	fmt.Println(i)
	// case 3:
	// 	fmt.Println(i)
	// case 4:
	// 	fmt.Println(i)
	// case 5:
	// 	fmt.Println(i)
	// default:
	// 	fmt.Println(i)
	// }

	// Multiple Switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Friday, time.Sunday:
	// 	fmt.Println("Weekend hai bhai maze marr")
	// default:
	// 	fmt.Println("Kaam Karo")
	// }

	// Type Switch
	whoamI := func(i interface{}) {
		// switch t := i.(type) {
		switch i.(type) {
		case int:
			fmt.Println("Its a integer")
		case string:
			fmt.Println("Its a string")
		case float32:
			fmt.Println("Its a float")
		case bool:
			// fmt.Println("other", t)
			fmt.Println("other")
		}

	}

	whoamI(12)
}

// Commmand to run this file -
// go run 10.go
