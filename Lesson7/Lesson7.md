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

## Errors

In Go, errors are a built-in type that represents an error condition. The `error` type is an interface that has a single method, `Error() string`, which returns a string representation of the error.
The `error` type is used to indicate that an operation has failed or that something unexpected has occurred. Functions that can return an error typically return a value of the type they are supposed to return, followed by an `error` value.
A nil `error` value indicates that no error occurred, while a non-nil `error` value indicates that an error occurred.

```go
package main
import (
    "errors"
    "fmt"
)
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```
