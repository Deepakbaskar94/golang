package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("the pointer value is:", ptr)

	number := 20
	var ptr = &number
	fmt.Println("the address in of number is:", ptr)
	fmt.Println("the value in that ptr address:", *ptr)

	*ptr = *ptr + 5
	//the value in number changes because *ptr actually referring to number variable value through address
	fmt.Println("the value in number is:", number)

	number = number + 5
	var ptr1 = &number
	fmt.Println("New address", ptr1)

	number = 35
	var ptr2 = &number
	fmt.Println("New address", ptr2)
}
