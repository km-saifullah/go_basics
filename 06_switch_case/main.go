package main

import (
	"fmt"
	"time"
)

func main() {
	num := 30
	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other", num)
	}

	// switch statement with multiple conditions
	switch time.Now().Weekday() {
	case time.Friday, time.Saturday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Workday")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		default:
			fmt.Println("Other", t)
		}
	}
	whoAmI("Go")
}
