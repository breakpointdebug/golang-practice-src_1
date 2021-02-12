package main

import (
	"fmt"
	"time"
)

func main() {
	yob := 2000

	for i := yob; i <= time.Now().Year(); i++ {
		fmt.Println(i)
	}

	fmt.Printf("you are %v years old today\n", time.Now().Year()-yob)
}
