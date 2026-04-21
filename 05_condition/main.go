package main

import "fmt"

func main() {

	// if block
	isActive := true
	if isActive == true {
		fmt.Println("User is active")
	}

	// if else block
	age := 15
	if age >= 18 {
		fmt.Println("Selected")
	} else {
		fmt.Println("Not Selected")
	}

	var role = "admin"
	var hasPermission = false

	if role == "admin" && hasPermission {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// else if block
	grade := 72

	if grade >= 80 && grade <= 100 {
		fmt.Println("Grade: A")
	} else if grade >= 70 && grade < 80 {
		fmt.Println("Grade: B")
	} else if grade >= 60 && grade < 70 {
		fmt.Println("Grade: C")
	} else if grade >= 33 && grade < 60 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}

	// variable declare in the if block
	if num := -2; num > 0 {
		fmt.Println("Positive number")
	} else {
		fmt.Println("Negative Number")
	}
}
