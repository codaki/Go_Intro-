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
