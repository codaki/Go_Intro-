package main

import "fmt"

type Animal interface {
	bark() string
}
type Dog struct {
	name string
}

func (f Dog) bark() string {
	message := "The dog is barking!!!"
	return message
}

func main() {
	var can Animal = Dog{name: "Felipe"}
	fmt.Println(can.bark())
}
