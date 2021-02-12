package main

import "fmt"

func main() {
	favSport := "hoy"
	switch favSport {
	case "hoy":
		fmt.Println("this is case1")
	case "hey":
		fmt.Println("this is case2")
	default:
		fmt.Println("this is default case")
	}
}
