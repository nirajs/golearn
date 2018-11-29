package main

import "fmt"

func main() {
	fmt.Println(greeting("niraj"))
	fmt.Println(getSum(3, 5))
}

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}
