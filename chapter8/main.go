package main

import (
	"fmt"
	"time"
)

// You cant declare variables outside of a function in Go

func main() {
	var roomNumber, floorNumber int = 153, 9
	fmt.Println(roomNumber, floorNumber)

	var psswd string = "password"
	fmt.Println(psswd)

	// short hand usage of the same from above

	roomNumber1, floorNumber1 := 153, 9
	fmt.Println(roomNumber1, floorNumber1)

	psswd1 := "password"
	fmt.Println(psswd1)

	// nil cannot be used with short hand declaration

	// an untyped constant has no type
	const pi = 3

	const occupancyLimit = 12

	var occupancyLimit1 uint8
	var occupancyLimit2 int64
	var occupancyLimit3 float32

	// assign our untyped const to an uint8 variable
	occupancyLimit1 = occupancyLimit
	// assign our untyped const to an int64 variable
	occupancyLimit2 = occupancyLimit
	// assign our untyped const to an float32 variable
	occupancyLimit3 = occupancyLimit

	fmt.Println(occupancyLimit1, occupancyLimit2, occupancyLimit3)

	const UKTimezoneName = "Europe/London"
	loc, err := time.LoadLocation(UKTimezoneName)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(loc)
	}
}
