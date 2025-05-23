package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	//Defining variables
	var name string = "Christopher"
	var age int = 70

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	//short hand declaration
	name1 := "Christopher"
	age1 := 40

	fmt.Println("Name in short hand: ", name1)
	fmt.Println("age in short hand: ", age1)
	//block declaration
	var (
		name2 string = "Christopher"
		age2  int    = 70
	)
	//multiple variable declaration
	var name3, age3 = "Christopher", 70
	var age4, age5 int = 70, 80
	fmt.Println("Name2:", name2)
	fmt.Println("Age2:", age2)
	fmt.Println("Name3:", name3)
	fmt.Println("Age3:", age3)
	fmt.Println("Age4:", age4)
	fmt.Println("Age5:", age5)

	//Constants
	const pi = 3.14

	fmt.Println("Constant", pi)
}
