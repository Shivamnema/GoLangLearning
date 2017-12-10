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
	}
}
