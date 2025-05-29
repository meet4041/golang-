package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/meet4041/podcast/auth"
	"github.com/meet4041/podcast/user"
)

// To create a package, we need to first initialize a module

// go mod init <package-name>
// go mod init github.com/meet4041/podcast

func main() {

	// Packages
	auth.Login("meet4041", "123456")
	session := auth.GetSession()
	fmt.Println(session)
	user := user.User{
		// Id:   1,
		Name: "meet4041",
		// Age:  22,
	}
	// fmt.Println(user)
	color.Green(user.Name)

}

// Commmand to run this file -
// go run 29.go
