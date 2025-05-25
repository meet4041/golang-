package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Id:", id)
}

func main() {

	// Wait Groups

	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
}

// Commmand to run this file -
// go run 25.go
