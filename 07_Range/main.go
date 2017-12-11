package main

import "fmt"

func main() {
	num := []int{3, 5, 5}
	sum := 0

	for _, num := range num {
		sum += num
	}
	fmt.Println("Sum", sum)
}
