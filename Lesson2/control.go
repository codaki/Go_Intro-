package main

import "fmt"

func main() {

	//If statement
	fmt.Println("If statement")
	var age int = 30
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are not an adult")
	}

	//Traditional for loo
	fmt.Println("Traditional for loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//For loop with range
	//range is used to iterate over a slice, array, map, or string
	fmt.Println("For loop with range")
	for i := range 5 {
		fmt.Println(i)
	}

	//Switch statement
	fmt.Println("Switch statement")
	value := 3
	switch value {
	case 1:
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("Value is 2")
	case 3:
		fmt.Println("Value is 3")
	default:
		fmt.Println("Value is not 1, 2 or 3")
	}

	dataStructures()

}
