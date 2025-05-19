package main

import "fmt"

// 'FUNCTION'
// go function can return multiple values
// func add(a, b int) int {
// 	return a + b
// }

func getLanguages() (string, string, int) {
	return "golang", "youtube", 1
}

func processIt(f func(a int) int) {
	fn(1)
}

func main() {
	// result := add(4, 5)
	// fmt.Println("Sum:", result)

	// lang1, lang2, lang3 := getLanguages()
	// fmt.Println(lang1, lang2, lang3)
	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)
	fn := func(a int) {
		return 33
	}

	processIt(fn)
}

// Commmand to run this file -
// go run 15.go
