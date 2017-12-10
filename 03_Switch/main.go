package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("It's One.")
	case 2:
		fmt.Println("It's Two.")
	case 3:
		fmt.Println("It's Three")
	}

	//
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend")
	case time.Monday:
		fmt.Println("It's Monday")
	}
	name := "shivam"
	switch name {
	case "shivam":
		fmt.Println("The Name is ", name)

	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
