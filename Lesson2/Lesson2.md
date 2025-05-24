# Lesson2

## If statement

```go
if condition {
    // code to be executed if condition is true
} else {
    // code to be executed if condition is false
}
```

## For loop

Traditional for loop syntax in Go is similar to C/C++.

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

Rnge loop is used to iterate over elements in a collection.

```go
for index, value := range collection {
    fmt.Println(index, value)
}
```

- The `range` keyword is used to iterate over elements in a collection, such as arrays, slices, maps, or strings.
- The `index` variable holds the index of the current element, and the `value` variable holds the value of the current element.
- The `range` loop automatically handles the iteration, so you don't need to manually increment the index.
- There are no while loops in Go, but the `for` loop can be used as a while loop.

For loop can also be used as a while loop.

```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

## Switch statement

```go
switch variable {
case value1:
    // code to be executed if variable == value1
case value2:
    // code to be executed if variable == value2
default:
    // code to be executed if variable doesn't match any case
}
```

The switch statements is similar to the ones in C/C++. The difference is that the switch statement in Go does not require a break statement to prevent fall-through behavior. Each case is treated as a separate block, and execution will not fall through to the next case unless explicitly specified using the `fallthrough` keyword.
( fallthrough is when you want to execute the next case even if the current case is true. )

## Multiple files management

- The main package is the entry point of a Go program. It contains the `main` function, which is executed when the program is run.

- Go allows you to split your code into multiple files for better organization and maintainability.

- You can use the `import` statement to include other packages or files in your code.

- Each file in a Go project must belong to a package. The package name is specified at the top of the file using the `package` keyword.

- It is not possible to have two main packages in a single project.
- If you have multiple files in the same package, they can share variables and functions without needing to import each other.
- Like Java and C#, Go uses the concept of packages to organize code. A package is a collection of related Go files that are compiled together. Each package has its own namespace, which helps avoid naming conflicts.

## Arrays

- An array is a fixed-size collection of elements of the same type.
- The size of an array is part of its type, so arrays of different sizes are considered different types.
- Arrays are value types, meaning that when you assign an array to another variable or pass it to a function, a copy of the array is made.
- There are no array methods in Go, but you can use the built-in `len` function to get the length of an array.
- There are several ways to create an array in Go:
  - Using the `var` keyword to declare an array with a specific size.
  - Using the `[...]` syntax to create an array with a size determined by the number of elements.
