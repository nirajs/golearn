package main

import "fmt"

func main() {
	// Array
	var fruitArr [2]string

	// Assing values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assing

	fruit := [2]string{"Apple", "Orange"}

	// Slices
	fruitSlice := []string{"AppleSlice", "OrangeSlice"}

	fmt.Println(fruitArr)
	fmt.Println(fruit)
	fmt.Println(fruitSlice, len(fruitSlice))
}
