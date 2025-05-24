package main

import (
	"fmt"
	"time"
)

// func task(id int) {
// 	fmt.Println("Id:", id)
// }

func main() {
	// Goroutines

	fmt.Println("hi")
	for i := 0; i <= 10; i++ {
		// go task(i)

		go func(i int) {
			fmt.Println(i)
		}(i) // for inline goroutines
	}
	time.Sleep(time.Second * 2)
}

// Commmand to run this file -
// go run 24.go
