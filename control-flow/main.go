package main

import "fmt"

func fnFor() {
	// for init; condition; post statement (after iteration) {
	// 	//...
	// }

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}

func fnForNested() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\tThe inner loop: %d\n", j)
		}
	}
}

func fnForBreakContinue() {
	// break (breaks out of the loop)
	// continue (go on top again of loop)

	x := 0
	for {
		if x == 1 {
			break
		}
		x++
	}

	// getting remainder using modulo operator
	y := 43 % 40
	fmt.Printf("%v\n\n", y)

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			continue
		}
	}
}

func fnForASCII() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%b\t%#x\t%#U\n", i, i, i, i)
	}
}

var yy int = 42

func fnIfElse() {
	if true {
		fmt.Println("predefined constant")
	}
	if !false {
		fmt.Println("negation")
	}

	if x := 42; x == yy {
		fmt.Println("initialization statement with logical statement")
	}
}

func fnLoopConditionalModulo() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func fnSwitch() {
	// add "fallthrough" keyword if you want to also print the next case
	// "default" keyword as the one catches if value on switch does not exist
	x := 5
	switch x {
	case 1, 5: // multiple cases
		fmt.Printf("case1 value was: %v\n", x)
	case 2:
		fmt.Printf("case2 value was: %v\n", x)
	case 3:
		fmt.Printf("case3 value was: %v\n", x)
		fallthrough
	default:
		fmt.Println("no value")
	}
}

func fnConditionalLogicOperator() {
	// AND: if anything is false then return as false, else true
	// OR: if anything is true then return as true, else false
	fmt.Printf("true && true\t%v\n", true && true)
	fmt.Printf("true && false\t%v\n", true && false)
	fmt.Printf("true || true\t%v\n", true || true)
	fmt.Printf("true || false\t%v\n", true || false)
	fmt.Printf("!true \t%v\n", !true)
}

func main() {
	// fnFor()
	// fnForNested()
	// fnForBreakContinue()
	// fnForASCII()
	// fnIfElse()
	// fnLoopConditionalModulo()
	// fnSwitch()
	fnConditionalLogicOperator()
}
