package main

import "fmt"

func main() {

	// Define map
	email := make(map[string]string)

	// declare map2
	email1 := map[string]string{"niraj": "sinha"}

	// Assing kv
	email["Bob"] = "bob@gmail.com"
	email["Sharon"] = "sharon@gmail.com"
	email["Bobe"] = "bob@gmail.com1"

	fmt.Println(email)
	fmt.Println(len(email))

	// delete

	delete(email, "Bobe")
	fmt.Println(email1)
}
