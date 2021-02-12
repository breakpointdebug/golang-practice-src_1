package main

import (
	"fmt"
)

func fnArray() {
	// you don't really use Arrays, you use slices
	var x [5]int
	fmt.Println(x)
	x[3] = 42 // zero based index
	fmt.Println(x)
	fmt.Println(len(x))
}

func fnSlices() {
	// x := type {values}  // composite literal
	x := []int{1, 2, 3, 4, 5, 6} // composite data type, value of same type
	fmt.Println(x)
}

func fnRangeVSSlices() {
	x := []string{"a", "b", "c", "d"}
	for i, v := range x { // can use _ to just ignore using the value returned by range
		fmt.Printf("for index: %v, value is: %v\n", i, v)
	}
}

func fnSlicingASlice() {
	x := []int{42, 69, 420, 1337}

	fmt.Print("value of array: \t\t\t\t\t{")
	for i, v := range x {
		if i == len(x)-1 {
			fmt.Printf("%v}\n", v)
			break
		}
		fmt.Printf("%v,", v)
	}

	// colon to slice a slice
	// [<start position at zero index> : <how many to take from that point>]

	fmt.Printf("just take everything (x[:]):\t\t\t\t%v\n", x[:])
	fmt.Printf("start at index 1, take everything (x[1:]):\t\t%v\n", x[1:])
	fmt.Printf("start at index 0, take until 2nd element (x[:2]):\t%v\n", x[:2])
	fmt.Printf("start at index 1, take until 3rd element (x[1:3]):\t%v\n", x[1:3])
}

func main() {
	// fnArray()
	// fnSlices()
	// fnRangeVSSlices()
	fnSlicingASlice()
}
