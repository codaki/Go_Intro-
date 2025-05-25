package operations

import (
	"math"
)

func Add(x, y float64) float64 {
	return x + y
}
func Sub(x, y float64) float64 {
	return x - y
}
func Mul(x, y float64) float64 {
	return x * y
}
func Div(x, y float64) float64 {
	return x / y
}
func Mod(x, y float64) float64 {
	return (math.Mod(x, y))
}
func Pow(x, y float64) float64 {
	return (math.Pow(x, y))
}
func Sqrt(x float64) float64 {
	return (math.Sqrt(x))
}
