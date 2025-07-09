package main

import (
	"fmt"
	"time"
)

func main() {
	NewCh := make(chan string)
	NewCh22 := make(chan string)

	//NOTE : Anonymous function which you forgot to call that's why it would throw error
	go func() {
		time.Sleep(20 * time.Millisecond)
		NewCh <- "Message from Channel 1"
	}()
	go func() {
		time.Sleep(30 * time.Millisecond)
		NewCh <- "Message from Channel 2"
	}()
	select {
	case msg1 := <-NewCh:
		fmt.Println("Received from Channel 1")
		fmt.Println(msg1)

	case msg2 := <-NewCh22:
		fmt.Println("Recieved from Channel -2", msg2)
		// default:
		// 	fmt.Println("No message Recieved yet")
	}

}

//NOTE : Go select only goes for the fastes reaction , leaving the slowest to dust and not using it
