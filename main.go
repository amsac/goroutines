package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42 // Sending a value to the channel
	}()

	value := <-ch      // Receiving a value from the channel
	fmt.Println(value) // Output: 42

	ch2 := make(chan int, 2) // Creates a buffered channel with capacity 2

	ch2 <- 1
	ch2 <- 2

	fmt.Println(<-ch2) // Output: 1
	fmt.Println(<-ch2) // Output: 2

	// In the below example, the unbuffered channel ch is able to hold three values because the sender goroutine and the
	// receiver goroutine are synchronized. The sender goroutine executes a loop and sends three values (1, 2, and 3) to the
	// channel ch. However, since there is a receiver waiting for each sent value, the sender blocks after each send operation
	//  until the receiver receives the value from the channel. This synchronization ensures that the channel does not become
	//  full and can hold all the values being sent.

	// Regarding the logs, it may appear as if the goroutines are executing synchronously because the receiver goroutine is
	//  waiting for values from the channel before printing them. The range loop used in the receiver goroutine receives values
	//  from the channel one at a time. It blocks until a value is available on the channel and then proceeds to print it. This
	//   behavior creates a synchronous flow between the sender and receiver goroutines, making it seem as if they execute sequentially.

	// In summary, although the channel ch is unbuffered, the sender goroutine blocks after each send operation until the receiver goroutine
	// receives the value, allowing the channel to hold multiple values temporarily. The range loop in the receiver goroutine ensures that
	// values are received one at a time, creating a synchronous behavior in the logs.

	ch3 := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch3 <- i
		}
		close(ch3) // Closing the channel after sending all values
	}()

	for value := range ch3 {
		fmt.Println(value)
	}

	// the select statement:

	// The select statement allows you to handle multiple channel operations simultaneously.
	//  It waits until one of the case statements is ready to proceed.

	ch4 := make(chan string)
	ch5 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch4 <- "Hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch5 <- "World"
	}()

	select {
	case msg1 := <-ch4:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch5:
		fmt.Println("Received:", msg2)
	}

}
