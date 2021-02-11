package main

import "fmt"

const (
	// year1 ...
	year1 = 2017 + iota
	// year2 ...
	year2
	// year3 ...
	year3
	// year4 ...
	year4
)

func main() {
	fmt.Printf("year1: \t%v\n", year1)
	fmt.Printf("year2: \t%v\n", year2)
	fmt.Printf("year3: \t%v\n", year3)
	fmt.Printf("year4: \t%v\n", year4)
}
