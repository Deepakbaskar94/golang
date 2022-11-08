package main

import "fmt"

func main() {
	fmt.Println("Welcome to array program")

	var fruitlist [4]string

	fruitlist[0] = "apple"
	fruitlist[1] = "banana"
	fruitlist[3] = "orange"
	fmt.Println("all fruit list data:", fruitlist)
	fmt.Println("all fruit list data:", len(fruitlist))

	var vegetablelist = [3]string{"onion", "brinjal", "potato"}
	fmt.Println("all vegetable list data", vegetablelist)
	fmt.Println("all vegetable list data", len(vegetablelist))

}
