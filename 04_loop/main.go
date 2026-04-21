package main

import "fmt"

func main() {
	// for loop
	for count := 0; count <= 10; count++ {
		fmt.Println(count)
	}

	// using range
	for num := range 5 {
		fmt.Println(num)
	}

	// infinite loop
	for {
		fmt.Println("Infinite")
	}
}
