package main

import "fmt"

type Person struct {
	name  string
	email string
}

func bio(newPerson Person) {
	fmt.Printf("name is %s\n", newPerson.name)
	fmt.Printf("email is %s\n", newPerson.email)
}

func main() {
	fmt.Println("Learning receiver function")
	newPerson := Person{
		name:  "sakil jomadder",
		email: "sakil.jomadder@example.com",
	}
	bio(newPerson)
}
