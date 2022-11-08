package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("welcome to user input program:")
	fmt.Println("Give feedback fro this code: ")
	reader := bufio.NewReader(os.Stdin)

	//comma ok syntax or comma error syntax
	input, _ := reader.ReadString('\n')

	fmt.Println("This is the score which you gave me: ", input)
	fmt.Printf("This is the type of the input which you gave me: %T  \n", input)

}
