package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// math package
	fmt.Println(math.Abs(-5))
	fmt.Println(math.Sqrt(144))
	fmt.Println(math.Pow(2, 4))
	fmt.Println(math.Ceil(3.25))
	fmt.Println(math.Floor(3.25))
	fmt.Println(math.Round(3.25))
	fmt.Println("Random number between 0-10 is", rand.Intn(10))

	// constant
	fmt.Println(math.Pi)
	fmt.Println(math.Phi)
	fmt.Println(math.E)

	// user input
	var num1, num2 float64

	fmt.Print("Enter the firtst number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	ans := num1 + num2
	fmt.Printf("The sum of %.2f + %.2f = %.3f\n", num1, num2, math.Round(ans))

	// date package
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	future := now.Add(24 * time.Hour)
	fmt.Println(future)
	fmt.Println(time.Now().Unix())

	t, _ := time.Parse("2006-01-02", "2026-04-21")
	fmt.Println(t)
}
