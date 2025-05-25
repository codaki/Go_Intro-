package main

import (
	"fmt"
)

func main() {

	map1 := make(map[string]int)
	fmt.Println("Enter your name:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Enter your age:")
	var age int
	// fmt.Scanf("%d", &age) // This is another way to read input
	fmt.Scanln(&age)

	fmt.Println("YOUR NAME IS:", name)
	fmt.Println("YOUR AGE IS:", age)
	fmt.Println("Lets check if your are able to vote")
	if age >= 18 {
		fmt.Println("You are able to vote")
	} else {
		fmt.Println("You are not able to vote")
	}
	map1[name] = age
	fmt.Printf("%v you are %d years old", name, age)

}
