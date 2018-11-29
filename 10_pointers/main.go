package main

import "fmt"

func main() {

	a := 5
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T %T\n", a, b)

	// Use to val from address
	fmt.Println(*b)

	// Change val with pointer

	*b = 10
	fmt.Println(a)

}
