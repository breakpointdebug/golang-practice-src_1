package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("this is true")
	case false:
		fmt.Println("this is false")
	default:
		fmt.Println("this is default")
	}
}
