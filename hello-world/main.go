package main

import (
	"fmt"
)

func main() {
	//fmt.Println("hey")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			// _ = throw away returns
			_, err := fmt.Println(i, 42, true)
			// fmt.Println(n)
			fmt.Println(err)
		}
	}
}
