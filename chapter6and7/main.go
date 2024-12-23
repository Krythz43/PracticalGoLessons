package main

import "fmt"

func main() {
	// 0x67D is 1661 in decimal
	n := 0x67D
	fmt.Printf("Number in decimal: %d", n)

	d := 1661
	fmt.Printf("Number in hex: %x\n", d) 

	fmt.Printf("\nHex with capital: %X\n", d) 

	fmt.Printf("Number in octal: %o\n", d)

	fmt.Printf("Octal with prefix %O\n", d)

	fmt.Printf("%U %U Golang",110001000010001,111001000110001)
}