# Lesson 5

## Function Closures

Go functions can be nested, and the inner function can access variables from the outer function's scope. This is known as a closure.

```go
func main() {
    outer := func(x int) func(int) int {
        return func(y int) int {
            return x + y
        }
    }

    addFive := outer(5)
    fmt.Println(addFive(10)) // Output: 15
}
```

Go functions may be closured over variables. This means that a function can capture and remember the values of variables from its surrounding context, even after that context has finished executing.

```go
package main
import "fmt"
func main() {
    x := 10
    increment := func() int {
        x++
        return x
    }

    fmt.Println(increment()) // 11
    fmt.Println(increment()) // 12
}
```
