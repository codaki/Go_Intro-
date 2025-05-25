package main

import (
	"fmt"
)

func main() {
	fmt.Println("THIS IS THE RETURN", deferExample())
	var valor int = 10
	var p *int = &valor
	fmt.Println("Este es el valor de p", p)
	fmt.Println("Ese es puntero de p", *p)
	fmt.Println("La persona es:", initializePerson("christopher", 1))

}

type Person struct {
	Name string
	Age  int
}

// Use of defer
func deferExample() int {
	fmt.Println("This goes first")
	for i := range 10 {
		defer fmt.Println(i)
	}
	return 1
}

func initializePerson(x string, y int) Person {
	structPerson := Person{x, y}
	return structPerson
}
