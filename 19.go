package main

import (
	"fmt"
	"time"
)

// Struct works as blue print
// We can createa multiple instance of a single struct

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // in nano-second precision
}

func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder

}

// Receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}
func (o order) changeAmount() float32 {
	return o.amount
}

func main() {

	// 'Struct'
	// myorder := order{
	// 	id:     "1",
	// 	amount: 100.45,
	// 	status: "received",
	// }
	// myorder.createdAt = time.Now()

	// myorder.status = "paid" // changing status value

	// myorder.changeStatus("confirmed")
	// fmt.Println(myorder.changeAmount())

	// fmt.Println(myorder.id)
	// fmt.Println(myorder.amount)
	// fmt.Println(myorder.status)
	// fmt.Println(myorder.createdAt)

	// fmt.Println("Order Struct:", myorder)

	myOrder := newOrder("1", 45.45, "Paid")
	fmt.Println(myOrder)

	// If struct needed to be used single time
	// Inline struct
	abc := struct {
		name   string
		isGood bool
	}{"Meet", true}
	fmt.Println(abc)
}

// Commmand to run this file -
// go run 19.go
