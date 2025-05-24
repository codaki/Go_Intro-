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

  ## Slice

- A slice is a dynamically-sized, flexible view into the elements of an array.
- Slices are more powerful and flexible than arrays, as they can grow and shrink in size.
- Helps to store values of the same type without needing to specify a fixed size.
- There are several ways to create a slice in Go:
  - Using the `make` function to create a slice with a specific length and capacity.
  - Using the `[]` syntax to create a slice from an existing array or another slice.
  - Using the `append` function to add elements to a slice.
  ```go
    // Creating a slice using the make function
    slice := make([]int, 0, 5) // length 0, capacity 5
    // Creating a slice from an array
    array := [5]int{1, 2, 3, 4, 5}
    slice := array[1:4] // slice from index 1 to 3 (not inclusive of 4)
    // Creating a slice from another slice
    slice1 := []int{1, 2, 3}
    slice2 := slice1[1:3] // slice from index 1 to 2 (not inclusive of 3)
  ```

The function `len` returns the length of a slice, and the function `cap` returns the capacity of a slice. The capacity is the maximum number of elements that can be stored in the slice without needing to allocate more memory.

```go
    slice := []int{1, 2, 3}
    fmt.Println(len(slice)) // Output: 3
    fmt.Println(cap(slice)) // Output: 3
```

### Slice with different lenght and capacity

```go
package main
import ("fmt")

func main() {
  arr1 := [6]int{10, 11, 12, 13, 14,15}
  myslice := arr1[2:4]

  fmt.Printf("myslice = %v\n", myslice)
  fmt.Printf("length = %d\n", len(myslice))
  fmt.Printf("capacity = %d\n", cap(myslice))
}
```

Output:

```myslice = [12 13]
length = 2
capacity = 4
```

-The length of a slice is the number of elements it contains, while the capacity is the maximum number of elements it can hold before needing to allocate more memory.

The slice starts from the third element of the array which has value 12 (remember that array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.). The slice can grow to the end of the array. This means that the capacity of the slice is 4.

If myslice started from element 0, the slice capacity would be 6.

### Append function

- The `append` function is used to add elements to a slice. It automatically resizes the slice if necessary.

```go
slice = append(slice, 4, 5, 6) // adds elements 4, 5, and 6 to the slice
```

- Append one slice to another slice

```go
slice1 := []int{1, 2, 3}
slice2 := []int{4, 5, 6}
slice3 := append(slice1, slice2...) // appends slice2 to slice1
```

- The `...` operator is used to expand the elements of a slice when passing it as an argument to a function. This allows you to pass a slice as individual elements.

### Copy function

- The `copy` function is used to copy elements from one slice to another.

```go
slice1 := []int{1, 2, 3, 4, 5}
slice2 := make([]int, len(slice1)-2) // creates a new slice with length 3
copy(slice2, slice1) // copies elements from slice1 to slice2
```

output:

```go
[1 2 3]
```
