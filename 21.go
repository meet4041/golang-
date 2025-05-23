package main

import "fmt"

type payment struct {
}

func (p payment) makepayment(amount float32) {
	razorpaymentGateway := razorpay{}
	razorpaymentGateway.pay(amount)
}

type razorpay struct {
}

func (p razorpay) pay(amount float32) {
	// logic for payment
	fmt.Println("Payment using razor pay:", amount)

}

func main() {

	// 'Interfaces'
}

// Commmand to run this file -
// go run 21.go
