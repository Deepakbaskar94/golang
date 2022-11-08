package main

import "fmt"

// variable name first letter capital so it will be a public variable
const Logintoken string = "asdfasdfasdf"

func main() {
	var username string = "Deepak"
	fmt.Println(username)
	fmt.Printf("variable type is: %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("variable type is: %T \n", isLoggedin)

	var smallval uint16 = 256
	fmt.Println(smallval)
	fmt.Printf("variable type is: %T \n", smallval)

	var bigval uint64 = 2560000000000000000
	fmt.Println(bigval)
	fmt.Printf("variable type is: %T \n", bigval)

	var intdata int = 2560000000000000000
	fmt.Println(intdata)
	fmt.Printf("variable type is: %T \n", intdata)

	var smallfloat float32 = 2560.567830453242353242
	fmt.Println(smallfloat)
	fmt.Printf("variable type is: %T \n", smallfloat)

	var bigfloat float64 = 2560.567830453242353242
	fmt.Println(bigfloat)
	fmt.Printf("variable type is: %T \n", bigfloat)

	//auto declaration of variable type
	var number = 123456
	fmt.Println(number)
	fmt.Printf("variable type is: %T \n", number)

	var data1 = "data"
	fmt.Println(data1)
	fmt.Printf("variable type is: %T \n", data1)

	value1 := 123
	fmt.Println(value1)
	fmt.Printf("variable type is: %T \n", value1)

	var check string
	fmt.Printf("variable type is: %T \n", check)

}
