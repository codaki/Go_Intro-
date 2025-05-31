# Lesson 6

## Interfaces

In go interfaces are a way to define a contract that types can implement. An interface specifies a set of methods that a type must have to be considered as implementing that interface.
An interface is defined using the `type` keyword followed by the interface name and the methods it requires. For example:

```go
type Animal interface {
    Speak() string
}
```

In this example, the `Animal` interface requires any type that implements it to have a `Speak` method that returns a string.

Interfaces are implemented implicitly in Go. This means that a type does not need to explicitly declare that it implements an interface. If a type has all the methods required by an interface, it is considered to implement that interface.
For example, we can define a `Dog` type that implements the `Animal` interface:

```go

type Dog struct{}
func (d Dog) Speak() string {
    return "Woof!"
}
```

Empty interfaces are a special case in Go. An empty interface, defined as `interface{}`, can hold values of any type. This is useful when you want to write functions that can accept any type of value.
For example, we can define a function that takes an empty interface as an argument:

```go
func PrintValue(v interface{}) {
    fmt.Println(v)
}
```

This function can accept any value, regardless of its type, and print it to the console.

## Example

Here is a complete example that demonstrates the use of interfaces in Go:

```go
package main
import (
    "fmt"
)
type Animal interface {
    Speak() string
}
type Dog struct{}
func (d Dog) Speak() string {
    return "Woof!"
}
type Cat struct{}
func (c Cat) Speak() string {
    return "Meow!"
}
func PrintAnimal(animal Animal) {
    fmt.Println(animal.Speak())
}
func main() {
    var dog Animal = Dog{}
    var cat Animal = Cat{}

    PrintAnimal(dog) // Output: Woof!
    PrintAnimal(cat) // Output: Meow!

    // Using empty interface
    PrintValue("Hello, World!")
    PrintValue(42)
}
func PrintValue(v interface{}) {
    fmt.Println(v)
}
```

In this example, we define an `Animal` interface with a `Speak` method. We then create two types, `Dog` and `Cat`, that implement the `Animal` interface. The `PrintAnimal` function takes an `Animal` interface as an argument and prints the result of the `Speak` method. Finally, we demonstrate the use of an empty interface with the `PrintValue` function, which can accept any type of value.
