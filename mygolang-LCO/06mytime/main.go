package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time code")
	currenttime := time.Now()
	fmt.Println(currenttime)
	fmt.Printf("Data Type: %T \n", currenttime)
	fmt.Println(currenttime.Format("02-01-2006 15:04:04 Monday"))

	//created date input
	createddate := time.Date(1994, time.November, 13, 16, 30, 10, 0, time.UTC)
	createddate1 := createddate.Format("02-01-2006 15:04:04 Monday")
	fmt.Println(createddate1)
}
