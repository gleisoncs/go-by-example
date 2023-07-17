package main

import "fmt"

func ForRange() {

	// create an array
	numbers := [5]string{"11", "22", "33", "44", "55"}

	// for loop with range
	for item := range numbers {
		fmt.Println(item)
	}

}
