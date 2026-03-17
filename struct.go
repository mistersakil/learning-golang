package main

import "fmt"

type User struct {
	name  string
	age   int
	email string
}

func main() {
	newUser := User{
		name:  "sakil jomadder",
		age:   70,
		email: "sakil.jomadder@example.com",
	}

	fmt.Println("Learning struct")
	fmt.Println("User name: ", newUser.name)
	fmt.Println("User age: ", newUser.age)
	fmt.Println("User email: ", newUser.email)

}
