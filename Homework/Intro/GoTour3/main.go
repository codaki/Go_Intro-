// Exercise: Fibonacci closure
// Let's have some fun with functions.

// Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	//Closure: The inner function "closes over" variables a and b, maintaining their state between calls
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	f := fibonacci()
	for range 10 {
		//State persistence: Each call to f() remembers the previous values and continues the sequence
		fmt.Println(f())
	}
}
