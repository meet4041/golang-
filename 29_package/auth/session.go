package auth

func extractSession() string {
	return "loggedin"
}

func GetSession() string {
	return extractSession()
}

// Commmand to run this file -
// go run session.go
