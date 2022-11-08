package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("code for myslices")

	var fruitlist = []string{"apple", "orange", "grapes"}
	fmt.Printf("type of fruitlist data %T \n", fruitlist)
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "pineapple", "banana")
	fmt.Println(fruitlist[1:3])

	score := make([]int, 4)
	score[0] = 111
	score[1] = 999
	score[2] = 222
	score[3] = 333
	// score[4] = 444

	score = append(score, 444, 555, 666)
	fmt.Println(score)
	fmt.Println(sort.IntsAreSorted(score))
	sort.Ints(score)
	fmt.Println(score)
	fmt.Println(sort.IntsAreSorted(score))

	//delete from slice
	var courses = []string{"pyton", "java", "c", "c++"}
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
