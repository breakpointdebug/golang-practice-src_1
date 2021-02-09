package main

import "fmt"

var y = 42 // declared with value infer int datatype from value

func main() {
	fmt.Println(`original value:`, y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)  // binary
	fmt.Printf("%x\n", y)  // hex
	fmt.Printf("%#x\n", y) // hex format
	y = 69420
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x", y, y, y)
}

// Printf = has formatting
// Println = add newline at the end
// Print = just print
// Sprint = print it into a string as output
//  '''  f
//  ''' ln
// Fprint = print to a file itself
//  '''  f
//  ''' ln
