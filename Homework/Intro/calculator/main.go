package main

import (
	"calculator/operations"
	"fmt"
)

func main() {
	var selection int

	var for_conditional int = 0
	for for_conditional == 0 {
		fmt.Println("Welcome to calculator!!")

		fmt.Println("1) Sum of two numbers")
		fmt.Println("2) Subtraction of two numbers")
		fmt.Println("3) Multiplication of two numbers")
		fmt.Println("4) Division of two numbers")
		fmt.Println("5) Modulus of two numbers")
		fmt.Println("6) Power of two numbers")
		fmt.Println("7) Square root of a number")
		fmt.Println("8) Exit")

		fmt.Println("Select an option:")
		fmt.Scanln(&selection)

		var valor1 float64
		var valor2 float64

		fmt.Println("Inserte el primer número:")
		fmt.Scanln(&valor1)

		if selection != 7 {
			fmt.Println("Inserte el segundo número:")
			fmt.Scanln(&valor2)
		}

		switch selection {
		case 1:
			fmt.Println("You selected sum!")
			fmt.Println("El resultado es :")
			fmt.Println(operations.Add(valor1, valor2))
		case 2:
			fmt.Println("You selected subtraction!")
			fmt.Println("El resultado es :")
			fmt.Println(operations.Sub(valor1, valor2))
		case 3:
			fmt.Println("You selected multiplication!")
			fmt.Println("El resultado es :")
			fmt.Println(operations.Mul(valor1, valor2))
		case 4:
			fmt.Println("You selected division!")
			fmt.Println("El resultado es :")
			fmt.Println(operations.Div(valor1, valor2))
		case 5:
			fmt.Println("You selected modulus!")
			fmt.Println("El resultado es :")
			fmt.Println(operations.Mod(valor1, valor2))
		case 6:
			fmt.Println("You selected power!")
			fmt.Println("El resultado es :")
			fmt.Println(operations.Pow(valor1, valor2))
		case 7:
			fmt.Println("You selected square root!")
			fmt.Println("El resultado es :")
			fmt.Println(operations.Sqrt(valor1))
		case 8:
			for_conditional = 1
		default:
			fmt.Println("Wrong selection!! Try again.")
		}
		// Add this to wait for Enter key press before continuing
		fmt.Println("\nPress Enter to continue...")
		fmt.Scanln()
	}
}
