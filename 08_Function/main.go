package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

//Same type of variables in parameters
func plusPlus(a, b, c int) int {
	return a + b + c
}

//multiple values return
func vals() (int, int) {
	return 3, 7
}

func main() {
	fmt.Println(add(234, 32))
	fmt.Println(plusPlus(1, 2, 3))

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// when you only want to accept the subset of the return values
	_, c := vals()
	fmt.Println(c)
}
