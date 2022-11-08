package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to maps or hash or key value pair")

	languages := make(map[string]string)

	languages["PY"] = "PYTHON"
	languages["JS"] = "JAVA SCRIPT"
	languages["SW"] = "SWIFT"
	languages["HTML"] = "Hyper text markup language"

	fmt.Println(languages)
	fmt.Println(languages["PY"])
	delete(languages, "HTML")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Print(key)
		fmt.Print(":")
		fmt.Println(value)
		fmt.Println("=======")
		fmt.Printf("for key data %v and for value data %v \n", key, value)

	}
}
