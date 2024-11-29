package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func sub(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("Welcome to mini calculator")
	fmt.Println("Enter two numbers: ")
	var a, b int
	fmt.Scanln(&a, &b)
	fmt.Println("Enter the operation you want to perform: ")
	fmt.Println("1. Addition")
	fmt.Println("2. Multiplication")
	fmt.Println("3. Division")
	fmt.Println("4. Subtraction")
	var operation int
	fmt.Scanln(&operation)

	switch operation {
	case 1:
		fmt.Println(add(a, b))
	case 2:
		fmt.Println(mul(a, b))
	case 3:
		fmt.Println(div(a, b))
	case 4:
		fmt.Println(sub(a, b))
	}

	fmt.Println("Thank you for using mini calculator")
}
