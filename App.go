package main

import (
	"fmt"
)

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Println("Enter your second number: ")
	fmt.Scanln(&num2)
	fmt.Print("Enter your desired operation: (+ - * / % ): ")
	fmt.Scanln(&operator)

	switch operator {

	case "+":

		fmt.Printf("%f %s %f = %f", num1, operator, num2, num1+num2)

	case "-":
		fmt.Printf("%f %s %f = %f", num1, operator, num2, num1-num2)

	case "*":
		fmt.Printf("%f %s %f = %f", num1, operator, num2, num1*num2)

	/*case "%":
	fmt.Printf("%f %s %f = %f", num1, operator, num2, mod(num1, num2))*/

	case "/":
		if num2 == 0 {
			fmt.Println("Can't divide by zero lol")
		} else {

			fmt.Printf("%f %s %f = %f", num1, operator, num2, num1/num2)
		}

	default:
		fmt.Println("Enter Valid Operator")
	}
}
