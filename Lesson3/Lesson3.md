# Lesson 3

## Defer

- The `defer` statement is used to ensure that a function call is performed later in the program execution.
- It is often used for cleanup tasks, such as closing files or releasing resources.
- The deferred function call is executed after the surrounding function returns, regardless of whether the function returns normally or due to a panic.

```go
func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
```

- Output:

```
Hello
World
```

Defer functions are executed in LIFO (Last In, First Out) order. This means that if multiple `defer` statements are present, the last one defined will be executed first. Uses a stack to keep track of the deferred functions.

```go
func main() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    defer fmt.Println("Third")
}
```

- Output:

```
Third
Second
First
```

## More types

### Pointers

- A pointer is a variable that stores the memory address of another variable.
- Pointers are declared using the `*` operator.
- The `&` operator is used to get the address of a variable.
- The `*` operator is also used to dereference a pointer, which means accessing the value stored at the address the pointer points to.

```go

    var x int = 10
    var p *int = &x // p is a pointer to x

```

```
    fmt.Println(*p) // Output: 10
    fmt.Println(p)  // Output: memory address of x
```

## Structs

- A struct is a composite data type that groups together variables (fields) under a single name.
- Structs are defined using the `type` keyword followed by the struct name and its fields.
- Structs can be used to create complex data types that represent real-world entities.
- They are similar to classes in other programming languages but do not support inheritance.

```go
type Person struct {
    Name string
    Age  int
}
func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p.Name) // Output: Alice
    fmt.Println(p.Age)  // Output: 30
}
```
