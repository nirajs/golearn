package main

import "fmt"
import "strconv"

// Define Person Struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// Birthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer reiver)
func (p *Person) getMarried(spouseLastname string) {
	if p.gender == "M" {
		return
	}
	p.lastName = spouseLastname

}

func main() {

	// Init person using struct
	person1 := Person{firstName: "Foo", lastName: "bar", city: "NY", gender: "F", age: 25}
	fmt.Println("Detail of person1 ", person1)

	// Alternative
	person2 := Person{"bar", "foo", "YN", "F", 25}

	fmt.Println("Detail of person2 ", person2)

	person1.hasBirthday()
	person1.hasBirthday()

	person1.getMarried("sandhu")
	fmt.Println(person1.greet())

}
