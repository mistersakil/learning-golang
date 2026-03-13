package main

import "fmt"

const basicSalary = 20000

func bank(bankName string) func(name string, additionalSalary float64) {
	const tada = 5000
	fmt.Println("Welcome to the", bankName)

	var createAccount = func(name string, additionalSalary float64) {
		fmt.Println("Account created for", name)
		fmt.Println("Gross Salary:", basicSalary+additionalSalary+tada)
	}

	return createAccount

}

func main() {
	var newBank = bank("Bangladesh Bank")
	newBank("Sakil jomadder", 5000.0)
	fmt.Println("================")
	newBank = bank("Dhaka Bank")
	newBank("Rasel jomadder", 10000.555)
}
