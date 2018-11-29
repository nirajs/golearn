package main

import "fmt"

func main() {

	ids := []int{1, 2, 4, 5, 8, 10}

	// loop
	for i, id := range ids {
		fmt.Printf("%d - Id %d \n", i, id)
	}

	// range with maps

	emails := map[string]string{"niraj": "sinha", "gulshan": "sinha"}

	for k, v := range emails {
		fmt.Printf("%s: %s \n", k, v)
	}

}
