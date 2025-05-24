package main

import "fmt"

func dataStructures() {
	//Arrays
	//Declaration
	var arr1 [1]int
	//Initialization
	arr1[0] = 1
	//Declaration and Initialization
	arr2 := [3]int{1, 2, 3}
	var arr3 = [3]int{1, 2, 3}
	//Array with inferred length
	arr4 := [...]int{1, 2, 4, 5, 6}
	//Array initialize specific index
	//also is not needed to initialize all elements
	//the rest of the elements will be initialized to 0
	arr5 := [...]int{0: 1, 1: 3, 2: 5}

	fmt.Println("Array1:", arr1)
	fmt.Println("Array2:", arr2)
	fmt.Println("Array3:", arr3)
	fmt.Println("Array4:", arr4)
	fmt.Println("Array5:", arr5)
	//Length of array
	// len() function is used to get the length of an array
	fmt.Println("Length of Array1:", len(arr5))

	//Dynamic array

}
