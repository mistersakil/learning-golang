package main

import (
	"fmt"
)

func add(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func main() {
	fmt.Println(add(5.00000005, 3.00000003))
}
