package main

import "fmt"

func main() {
	fmt.Println("welcome to loops")

	// days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	// FOR LOOP
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("the index is %v and value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("the index is %v and value is %v\n", index, day)
	// }

	// WHILE LOOP
	roughvalue := 1
	for roughvalue < 10 {
		// goto code
		if roughvalue == 8 {
			goto jumpcode
		}

		// break code
		if roughvalue == 5 {
			// break
			// continue code
			roughvalue++
			continue
		}
		fmt.Println("the rough value is ", roughvalue)
		roughvalue++
	}

jumpcode:
	fmt.Println("this is jump code")
}
