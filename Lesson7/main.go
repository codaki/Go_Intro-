package main

import "fmt"

type Animal interface {
	bark() string
}
type Dog struct {
	name string
}
type Cat struct {
	name string
}

// This method means type Dog implements the interface Animal,
// but we don't need to explicitly declare that it does so.
func do(i interface{}) {
	switch v := i.(type) {
	case Dog:
		fmt.Println("Dog:", v.name)
	case Cat:
		fmt.Println("Cat:", v.name)
	default:
		fmt.Println("Unknown type")
	}
}
func (f Dog) bark() string {
	message := "The dog is barking!!!"
	return message
}

func main() {
	var can Animal = Dog{name: "Felipe"}
	fmt.Println(can.bark())

	//var cat Animal = Cat{name: "Josefino"} is not allowed
	cat1 := Cat{name: "Josefino"}
	fmt.Println(cat1.name)

	do(cat1)
	do("Hello, World!")
}
