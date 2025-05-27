# Lesson 4

## Nil Slices and Maps

The zero value of a slice is `nil`, which means it has no elements and no underlying array. A nil slice behaves like an empty slice, but it is not the same as an empty slice.

```go
var mySlice []int // nil slice
fmt.Println(mySlice == nil) // true
```

```go
var myMap map[string]int // nil map
fmt.Println(myMap == nil) // true
```

## Slices of Slices

A slice can contain other slices, allowing you to create a two-dimensional slice (or more dimensions).

```go
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}
fmt.Println(matrix[0]) // [1 2 3]
fmt.Println(matrix[1][2]) // 6
```

## Range

The `range` keyword is used to iterate over elements in a slice or map. It returns both the index (or key) and the value.

```go
for index, value := range mySlice {
    fmt.Println(index, value)
}
for key, value := range myMap {
    fmt.Println(key, value)
}
```
