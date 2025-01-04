package main

import (
	"Math/rand"
	"fmt"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Hello, World!")

	itr := 0

	for itr < 10 {
		fmt.Println("itr: ", itr)
		itr++
	}

	x, y := swap(5, 10)
	fmt.Println("x: ", x, "y: ", y)

	x, y = swap(x, y)
	fmt.Println("x: ", x, "y: ", y)

	swapWithReference(&x, &y)
	fmt.Println("x: ", x, "y: ", y)
}

func swap(x int, y int) (int, int) {
	return y, x
}

func swapWithReference(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
