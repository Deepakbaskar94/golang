package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	// assigning values inside a class
	ajay := User{"ajay dev", 25, "engineer"}
	fmt.Println(ajay)
	fmt.Printf("The data for ajay is %+v \n", ajay)
	fmt.Printf("Ajay is an %+v \n", ajay.Job)
}

// struct is like a class (no inheritance, super class, parent or child concept in go)
// all start letter should be capital to make it public
type User struct {
	Name string
	Age  int
	Job  string
}
