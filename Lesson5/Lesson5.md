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
