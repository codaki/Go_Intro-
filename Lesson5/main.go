package main

import (
	"fmt"
	"math"
)

// TODO: Implement a function called "counter" that:
// 1. Returns a function that increments and returns a counter
// 2. Each call should return the next number in sequence (1, 2, 3, ...)
// 3. Multiple counters should be independent

func counter() func() int {
	a := 1
	return func() int {
		result := a
		a++
		return result
	}
}

// TODO: Implement a function called "accumulator" that:
// 1. Takes an initial value as parameter
// 2. Returns a function that adds a value to the running total
// 3. The returned function should return the new total after adding

func accumulator(initial int) func(int) int {
	total := initial
	return func(x int) int {
		total += x
		return total
	}
}

type Exponent struct {
	y float64
	x float64
}

func (s Exponent) exec() float64 {
	return math.Pow(s.x, s.y)
}

func main() {
	// Test your implementation
	c1 := counter()
	c2 := counter()

	fmt.Println(c1()) // Should print: 1
	fmt.Println(c1()) // Should print: 2
	fmt.Println(c2()) // Should print: 1 (independent counter)
	fmt.Println(c1()) // Should print: 3
	fmt.Println(c2()) // Should print: 2

	// Test your implementation
	acc1 := accumulator(10) // Start with 10
	acc2 := accumulator(0)  // Start with 0

	fmt.Println(acc1(5))  // Should print: 15 (10 + 5)
	fmt.Println(acc1(3))  // Should print: 18 (15 + 3)
	fmt.Println(acc2(7))  // Should print: 7 (0 + 7)
	fmt.Println(acc1(2))  // Should print: 20 (18 + 2)
	fmt.Println(acc2(-3)) // Should print: 4 (7 + (-3))

	metodo := Exponent{x: 10, y: 2}
	fmt.Println(metodo.exec())

}
