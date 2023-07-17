package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string)

	go func() { messages <- ("ping - ") }()

	fmt.Println("waiting 5 seconds")
	time.Sleep(5 * time.Second)

	msg := <-messages
	fmt.Println(msg)
}
