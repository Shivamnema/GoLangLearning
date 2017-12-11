package main

import "fmt"

//Variadic functions can be called with any number of trailing arguments.
// For example, fmt.Println is a common variadic function.
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
}

func main() {
	sum(1, 3)
	sum(3, 3, 3)
}
