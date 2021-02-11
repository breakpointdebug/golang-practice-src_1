package main

import (
	"fmt"
	"runtime"
)

// https://golang.org/ref/spec

func fnBoolean() {
	a := 7
	b := 42

	fmt.Println(a == b)
}

func fnNumeric() {
	a := 42
	b := 42.42
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}

func fnMisc() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}

func fnString() {
	s := `"hey world ñÑàØµ ■²"` // immutable, but can assign a new value
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) // convert into a slice of bytes
	fmt.Println(bs)
	fmt.Printf("%T\n", bs) // byte is alias for uint8

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	fmt.Println("\n")

	for index, value := range s {
		fmt.Printf("at index position %d we have hex %#x\n", index, value) // display in hex value
	}
}

func fnNumeralSystems() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("Binary: %b\n", n)
	fmt.Printf("lower case hex: %x\n", n)
	fmt.Printf("hex in proper format: %#x\n", n)

}

func main() {
	// fnBoolean()
	// fnNumeric()
	// fnMisc()
	// fnString()
	// fnNumeralSystems()
}
