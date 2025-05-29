package auth

import "fmt"

func Login(username string, password string) {
	fmt.Println("Login successful:", username, password)
}

// Commmand to run this file -
// go run credientials.go
