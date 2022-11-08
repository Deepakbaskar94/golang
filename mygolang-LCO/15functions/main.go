package main

import "fmt"

func main() {
	fmt.Println("welcome to Functions program")
	greeter()
	result := adder(3, 5)
	fmt.Println("the result is: ", result)

	proresult, mymessage := proadder(2, 5, 7, 6, 3, 1, 4)
	fmt.Println("the pro result is:", proresult)
	fmt.Println("this is my message:", mymessage)

}

func greeter() {
	fmt.Println("Namaste")
}

func adder(valueone int, valuetwo int) int {
	resultone := valueone + valuetwo
	return resultone
}

func proadder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		fmt.Println("print value:", value)
		total += value
	}
	return total, "this is my message"

}
