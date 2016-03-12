package main

import (
	"fmt"
	"time"
)

func main() {
	//for loops and if
	for z := 0; z <= 10; z++ {
		fmt.Println(z)
		if z%2==0 {
			fmt.Println(z, " is an even number.")
		}else{
			fmt.Println(z, " is an odd number.")
		}
	}
//switch statement and time function
	time := time.Now()
	switch {
	case time.Hour() < 12:
		fmt.Println("Good Morning!")
	default:
		fmt.Println("Good Afternoon!")
	}

	//defer statements
	fmt.Println("\nLast In ----> First Out\n")
	for x := 0; x < 5; x++{
		defer fmt.Println(x, "popped!")
		fmt.Println(x, "deffered")
	}
	fmt.Println("\n")

}
