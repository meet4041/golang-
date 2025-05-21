package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone int
}

// composition
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // in nano-second precision
	customer            // this is called struct emdedding
}

func main() {

	// 'Struct Enbedding'

	newCustomer := customer{
		name:  "Meet",
		phone: 1212,
	}

	newOrder := order{
		id:       "1",
		amount:   300.02,
		status:   "received",
		customer: newCustomer,
	}

	newOrder.customer.name = "Meet Gandhi"
	newOrder.createdAt = time.Now()
	fmt.Println(newOrder)
}

// Commmand to run this file -
// go run 20.go
