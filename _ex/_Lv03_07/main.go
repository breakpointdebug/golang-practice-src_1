package main

import (
	"fmt"
	"time"
)

func main() {
	rng := time.Now().Second() % 2
	if rng > 0 {
		fmt.Println("shut up")
	} else {
		fmt.Println("ok, talk")
	}
}
