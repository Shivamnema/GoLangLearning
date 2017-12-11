package main

import "fmt"

//maps
func main() {
	fmt.Println("Hello")
	m := make(map[string]int)

	//initialise the map
	m["k1"] = 123
	m["k2"] = 123
	m["k3"] = 123
	
	fmt.Println(m)
	//length of map
	fmt.Println(len(m))

	//del any map using key
	delete(m, "k2")
	fmt.Println("map:", m)
	
	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

}
