package main

import "fmt"

func main() {
	var a [4]int
	a[3] = 324
	fmt.Println(a[3])
	fmt.Println(len(a))

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
