package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to this code")
	fmt.Println("Give rating for this code:")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	value, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Given input is:", value+1)
		//fmt.Println(err1)
	}

}
