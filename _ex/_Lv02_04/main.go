package main

import "fmt"

func main() {
	X := 42
	fmt.Printf("decimal: %d, binary: %b, hex: %#x\n", X, X, X)

	Y := (X << 1) // shift bit once to the left
	fmt.Printf("decimal: %d, binary: %b, hex: %#x\n", Y, Y, Y)
}
