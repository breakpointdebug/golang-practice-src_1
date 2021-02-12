package main

import (
	"fmt"
	"time"
)

func main() {
	yob := 2000

	for {
		if yob <= time.Now().Year() {
			fmt.Printf("%v\n", yob)
			yob++
			continue
		} else {
			break
		}
	}
}
