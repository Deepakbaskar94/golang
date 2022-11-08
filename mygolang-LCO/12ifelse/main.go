package main

import "fmt"

func main() {
	fmt.Println("Welcome to if else program")

	// type one
	logincount := 10
	var result string
	if logincount < 10 {
		result = "less than 10"
	} else if logincount > 10 {
		result = "greater than 10"
	} else {
		result = "equal to 10"
	}
	fmt.Println(result)

	// type two
	if 9%2 == 0 {
		fmt.Println("Number is EVEN")
	} else {
		fmt.Println("Number is ODD")
	}

	// type three
	if num := 3; num > 5 {
		fmt.Println("the number is greater")
	} else {
		fmt.Println("the number is smaller")
	}
}
