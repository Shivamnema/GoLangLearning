package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Value of X: ", x)
	foo()
}

func foo() {

	//This code does not compile as it's outside the scope of x.
	fmt.Println("X in Foo: ", x)
}
