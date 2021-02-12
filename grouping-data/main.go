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

func fnAppendValueToSlice() {
	x := []int{1, 2, 3}
	// append() is a special built in function just like len()
	// it is found at "builtin" package

	fmt.Println(append(x, 4, 5, 6)) // append at end of slice

	y := []int{111, 222, 333}
	fmt.Println(append(x, y...)) // append slice to a slice (should have the same data type)
}

func fnDeleteFromSlice() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x)
	// slice the first 2 elements
	// slice the elements starting from index 4
	// then append (combine), ignoring element: 3 & 4
	fmt.Println(append(x[:2], x[4:]...)) // deleted 3 and 4
}

func fnSliceMake() {
	// slices are dynamic, they change size automatically

	// if you know size of the array, use make()

	x := make([]int, 7, 12) // make(<type>, <length, indexes>, <capacity according to it's type>)

	fmt.Printf("slice: \t\t%v\n", x)
	fmt.Printf("length: \t%v\n", len(x))   // 7
	fmt.Printf("capacity: \t%v\n", cap(x)) // 12

	x = append(x, 300) // extend length automatically
	fmt.Printf("slice: \t\t%v\n", x)
	fmt.Printf("length: \t%v\n", len(x))   // 8
	fmt.Printf("capacity: \t%v\n", cap(x)) // 12

	x = append(x, 301)
	x = append(x, 302)
	x = append(x, 303)
	x = append(x, 304)
	x = append(x, 305)
	fmt.Printf("slice: \t\t%v\n", x)
	fmt.Printf("length: \t%v\n", len(x))   // 13
	fmt.Printf("capacity: \t%v\n", cap(x)) // 24
}

func fnMultiDimensionalSlice() {
	x := []string{"A", "B", "C"}
	fmt.Println(x)

	y := []string{"D", "E", "F"}
	fmt.Println(y)

	z := [][]string{x, y} // multi-dimensional slice
	fmt.Println(z)

	a := append(x, y...)
	fmt.Println(a)

	b := append(z, z...)
	fmt.Println(b)
}

func main() {
	// fnArray()
	// fnSlices()
	// fnRangeVSSlices()
	// fnSlicingASlice()
	// fnAppendValueToSlice()
	// fnDeleteFromSlice()
	// fnSliceMake()
	fnMultiDimensionalSlice()
}
