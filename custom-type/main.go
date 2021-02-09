package main

import "fmt"

var a int

type hotdog int // type hotdog, underlying type is int...

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// cannot do this: a = b, because they are of different types
	// you can't dynamically change a type because go is static
}
