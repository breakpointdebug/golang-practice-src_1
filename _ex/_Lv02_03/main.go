package main

import "fmt"

const (
	// A ...
	A = 42
	// B ...
	B string = "42"
)

func main() {
	fmt.Printf(`Untyped Constant "A" has value of: <%v> and type of: %T`, A, A)
	fmt.Println("")
	fmt.Printf(`Typed Constant "B" has value of: <%v> and type of: %T`, B, B)
	fmt.Println("")
}
