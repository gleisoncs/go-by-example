package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	var a = "teste"

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}(a)

	a = "agora vai"

	time.Sleep(time.Second)
	fmt.Println("done")
}
