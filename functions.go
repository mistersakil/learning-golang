package main

import "fmt"

func add(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func getNumbers(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 * num2
}

func getInputs() (int, int) {
	num1 := 0
	num2 := 0
	fmt.Println("Enter first number:")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number:")
	fmt.Scanln(&num2)
	return num1, num2
}

func main() {
	num1, num2 := getInputs()
	fmt.Println("The sum of numbers is:", add(float64(num1), float64(num2)))
	// fmt.Println(getNumbers(10, 20))
	// fmt.Println(add(5.00000005, 3.00000003))
}
