package main

import "fmt"

func main() {
	var name, op string
	var num1, num2 float32
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
	fmt.Println("Enter First number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter Second number: ")
	fmt.Scanln(&num2)
	fmt.Println("Enter the operator: ")
	fmt.Scanln(&op)
	fmt.Println("The result is:\t ", calculator(num1, num2, op) != -9999, "And We are", "Avoiding division by zero")

}

func calculator(num1 float32, num2 float32, op string) float32 {
	var result float32
	switch op {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 != 0 {
			return num1 / num2
		} else {
			fmt.Println("Division by zero is not possible")
			return -9999
		}
	default:
		fmt.Println("Invalid operator")
	}
	return result
}
