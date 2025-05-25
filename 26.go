package main

import "fmt"

// 2. send
// func numProcess(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing numbers...", num)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// 3. receive
// func sum(result chan int, num1 int, num2 int) {
// 	ans := num1 + num2
// 	result <- ans
// }

// 4. goroutine synchronization
//	func task(done chan bool) {
//		defer func() { done <- true }()
//		fmt.Println("processing task...")
//	}

// 5. Email Receiver
// func emailReceiver(emailChannel <-chan string, done chan<- bool) {
// 	defer func() { done <- true }()
// 	for email := range emailChannel {
// 		fmt.Println("Received email:", email)
// 		time.Sleep(time.Second * 1)
// 	}
// }

func main() {

	// Channels

	// 6.
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "Meet Gandhi"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received from chan2", chan2Val)
		}
	}

	// 5.
	// emailChannel := make(chan string, 100)
	// done := make(chan bool)
	// go emailReceiver(emailChannel, done)
	// for i := 1; i <= 5; i++ {
	// 	emailChannel <- fmt.Sprintf("%d@gmail.com", i)
	// }
	// fmt.Println("All emails sent")
	// // this is imp.
	// close(emailChannel)
	// <-done

	// 4.
	// done := make(chan bool)
	// go task(done)
	// <-done // blocking

	// 3.
	// result := make(chan int)
	// go sum(result, 10, 20)
	// ans := <-result
	// fmt.Println(ans)

	// 2.
	// numChan := make(chan int)
	// go numProcess(numChan)
	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// 1.
	// messageChan := make(chan string)
	// messageChan <- "Meet Gandhi" // blocking
	// msg := <-messageChan
	// fmt.Println(msg)

}

// Commmand to run this file -
// go run 26.go
