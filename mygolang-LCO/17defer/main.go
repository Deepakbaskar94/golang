package main

import "fmt"

func main() {

	defer fmt.Println("this is one")
	defer fmt.Println("this is Two")
	defer fmt.Println("this is Three")
	fmt.Println("this is Defer")
	mydefer()

}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}
