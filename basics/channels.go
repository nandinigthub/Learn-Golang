package basics

import (
	"fmt"
)

// send func
// func processNum(numChan chan int) {
// 	fmt.Println("process nnum", <-numChan)
// 	time.Sleep(time.Second)

// }

// // receiver func
// func sumnum(numChan chan int, a int, b int) {
// 	sumRes := a + b
// 	numChan <- sumRes
// }

// func emailSender(emailchan chan string, done chan bool) {
// 	defer func() { done <- true }()

// 	for email := range emailchan {

// 		fmt.Println("sending email to: ", email)
// 		// time.Sleep(time.Second)

// 	}
// }

func Channel1() {

	// emailchan := make(chan string, 10)
	// done := make(chan bool)

	// go emailSender(emailchan, done)

	// for i := 0; i < 10; i++ {
	// 	emailchan <- fmt.Sprintf("%d@mail.com", i)
	// }

	// fmt.Println("done sending emails address")

	// close(emailchan) // this is imp
	// <-done           // deadlock becoz it is a blocking we need to close channel

	// chan1 := make(chan int)

	// go func(a, b int) {
	// 	fmt.Println(a + b)
	// }(4, 5)
	// time.Sleep(time.Second * 2)

	// call := <-chan1
	// fmt.Println(call)

	// // receiver
	// go sumnum(chan1, 4, 51)
	// res := <-chan1
	// fmt.Println(res)

	// // sending
	// go processNum(chan1)
	// chan1 <- 5

	// // basic reciving and sending cause deadlock both are blocking on way connection

	// receive := make(chan int)
	// receive <- 1 // receiveing data in channel

	// send := <-receive // sending data from channel
	// fmt.Println(send)

	// for working with multiple channels
	// we use select case and inline goroutine func

	chan1 := make(chan string)
	chan2 := make(chan int)

	go func() {
		chan1 <- "hii"
	}()

	go func() {
		chan2 <- 5
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("receiving data from chan1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("receiving data from chan2", chan2Val)
		}

	}

}
