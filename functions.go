package main

import (
	"fmt"
)

func add(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func getNumbers(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 * num2
}

func main() {
	fmt.Println(add(5.00000005, 3.00000003))
	fmt.Println(getNumbers(10, 20))
}
