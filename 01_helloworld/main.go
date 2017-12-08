package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println("Add", 1+234)
	fmt.Println(!(true || false))

	// var t = 2342.234
	first := "shivam"
	fmt.Println(first)

	const pi = 3.14
	fmt.Println(pi)
	vi := pi + 23
	fmt.Println(vi)

	//Different ways of for loops in Go.
	// var i = 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for {
	// 	fmt.Println("loop")
	// 	break
	// }
	// for j := 7; j <= 9; j++ {
	// 	fmt.Println(j)
	// }

	//If/else condition in Go

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	i := 10
	for i <= 15 {
		if i == 10 {
			fmt.Println("hurry")
		}
		i++
	}
}
