package main

import "fmt"

func main() {
	num1 := 42
	num2 := 69
	fmt.Printf("== (num1=%d,num2=%d) result: %v\n", num1, num2, (num1 == num2))
	fmt.Printf("<= (num1=%d,num2=%d) result: %v\n", num1, num2, (num1 <= num2))
	fmt.Printf(">= (num1=%d,num2=%d) result: %v\n", num1, num2, (num1 >= num2))
	fmt.Printf("!= (num1=%d,num2=%d) result: %v\n", num1, num2, (num1 != num2))
	fmt.Printf("< (num1=%d,num2=%d) result: %v\n", num1, num2, (num1 < num2))
	fmt.Printf("> (num1=%d,num2=%d) result: %v\n", num1, num2, (num1 > num2))
}
