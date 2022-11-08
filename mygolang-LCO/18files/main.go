package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files")
	content := "this is the content inside the file"
	file, err := os.Create("./mytextfile.txt")
	errorchecknil(err)
	// if err != nil {
	// 	panic(err)
	// }

	length, err := io.WriteString(file, content)
	errorchecknil(err)

	fmt.Println("this is the length:", length)
	defer file.Close()
	readfile("./mytextfile.txt")

}

func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	errorchecknil(err)
	fmt.Println("this is databyte \n", string(databyte))
}

func errorchecknil(err error) {
	if err != nil {
		panic(err)
	}
}
