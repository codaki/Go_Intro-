# Lesson 1

## What is Go?

- Go was developed by Google.
- It is a compiled language.
- It is statically typed.
- It is a concurrent language( good at handling many things at once).
- It is used for high performance applications.
- It is not an object oriented language, but it has some features of OOP.

## Uses

- Web development
- Cloud computing
- cross platform applications

## Syntax for execution

In order to run a go program run the next command in the terminal.

```go
go run .
```

## Hello World

```go
package main
import "fmt"
func main() {
    fmt.Println("Hello, World!")
}
```

- In this part **package main** helps us to determine that it is a main package.
- the fmt package is imported in order to use Println function.

## Variables

There are two ways of defining a variable, regular and short hand declaration.

- Variables are case sensitive.

### Regular declaration

```go
var name string = "Christopher"
var age int = 23
```

### Short hand declaration

```go
name := "Christopher"
age := 23
```

## Data Types

There are **five** main data types in Go.

- **int**: for integers
- **float**: for floating point numbers
- **string**: for strings
- **bool**: for boolean values
- **complex**: for complex numbers
