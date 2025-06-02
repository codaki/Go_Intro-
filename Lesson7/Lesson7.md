# Lesson 7

## Type Switches

In Go, a type switch is a way to perform different actions based on the dynamic type of an interface value. It allows you to check the type of an interface value at runtime and execute code based on that type.
A type switch is similar to a regular switch statement, but it checks the type of the value rather than its value. The syntax for a type switch is as follows:

```go
switch v := value.(type) {
case Type1:
    // Handle Type1
case Type2:
    // Handle Type2
default:
    // Handle other types
}
```
