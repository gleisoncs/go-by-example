package main

import (
	"fmt"
)

func task1(str chan string) {
	fmt.Println(str)
}

func main() {
	done := make(chan string)
	go task1(done)
	done <- "world"
	fmt.Println(done)
}
