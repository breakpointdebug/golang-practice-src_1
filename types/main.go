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

	fmt.Println("")

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

// constant (typed vs untyped constants)
const (
	A         = 42
	B float64 = 42.42
	C string  = "James Bond"
)

// D ...
const D = 69

// E ...
const E = 420

const (
	// AI ...
	AI = iota
	// BI ...
	BI
	// CI ...
	CI
)

const (
	// DI ...
	DI = iota
	// EI ...
	EI
	// FI ...
	FI
)

func fnIota() {
	fmt.Println(AI)
	fmt.Println(BI)
	fmt.Println(CI)
	fmt.Println(DI)
	fmt.Println(EI)
	fmt.Println(FI)

}

func fnBitShifting() {
	X := 4
	fmt.Printf("decimal: %d\t\tbinary: %b\n", X, X)

	Y := X << 1 // shift bit value of X once
	fmt.Printf("decimal: %d\t\tbinary: %b", Y, Y)
}

const (
	_  = iota // do not use the 1st iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func fnSizeUsingIotaAndBitShifting() {
	// kb := 1024
	// mb := 1024 * kb
	// gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}

func main() {
	// fnBoolean()
	// fnNumeric()
	// fnMisc()
	// fnString()
	// fnNumeralSystems()
	// fnBitshifting()
	fnSizeUsingIotaAndBitShifting()
}
