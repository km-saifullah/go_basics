package main

import "fmt"

func main() {

	// variable name along with type
	var name string = "Go"
	fmt.Println(name)

	// variable name without the type
	var user = "Admin"
	fmt.Println(user)

	// shorthand to declare a variable
	language := "Go Language"
	fmt.Println(language)
	fmt.Printf("Type of Language is: %T\n", language)

	isActive := false
	fmt.Println(isActive)

	// simple mathematical calculation
	price := 150.0
	discount := 2.5
	fmt.Println(price - (price*discount)/100)
	fmt.Printf("Type of Price is: %T\n", price)

	// zero values
	var digit int
	var str string
	var isAuthenticated bool
	fmt.Println(digit, str, isAuthenticated) // digit = 0, str = "" , isAuthenticated = false

	// constants
	const pi = 3.1416
	const username = "user_xyz"
	fmt.Println(pi, username)

	// multiple variable declaration
	var num1, num2, num3 int = 10, 20, 30
	fmt.Println(num1, num2, num3)

	randomNum, randomStr := 55, "Go"
	fmt.Println(randomNum, randomStr)

	// group variable declaration
	var (
		title      = "Go Basics Course"
		duration   = 12
		isUploaded = true
	)
	fmt.Printf("Course Name: %s, Duration: %d hours, Uploaded: %t\n", title, duration, isUploaded)
}
