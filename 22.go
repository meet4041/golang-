package main

import "fmt"

// We implemt emums in GO using const

// type OrderStatus int

// const (
// 	Ordered OrderStatus = iota
// 	Packed
// 	Prpared
// 	Dispatched
// 	Delivered
// )

type OrderStatus string

const (
	Ordered    OrderStatus = "Ordered"
	Packed                 = "Packed"
	Prpared                = "Prepared"
	Dispatched             = "Dispatched"
	Delivered              = "Delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to:", "'", status, "'")
}

func main() {

	// Enums

	changeOrderStatus(Prpared)
}

// Commmand to run this file -
// go run 22.go
