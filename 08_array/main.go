package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 10
	fmt.Println(arr)

	var nums [4]int = [4]int{10, 20, 30, 40}
	fmt.Println(nums)

	numbers := [5]int{1, 2, 3, 4, 5}
	numbers[2] = 20 // update the value
	fmt.Println(numbers)
	fmt.Println(len(numbers)) // check the length of the array
	fmt.Println(numbers[2])

	even := [...]int{2, 4, 6, 8}
	fmt.Println(even)
	fmt.Println(len(even))

	// loop through an array
	for i, v := range nums {
		fmt.Println("Index:", i, "Value:", v)
	}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	names := [2]string{"Go", "Java"}
	fmt.Println(names)

	// 2d array
	matrix := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(matrix[0][1])
}
