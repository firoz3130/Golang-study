package main

/*
-->In Go, the := operator is a shorthand for declaring and initializing a variable. It is equivalent to declaring a variable with a specific type and then assigning a value to it. For example, the following two statements are equivalent:
	var name string
	name = "John"
	name := "John"
-->strings.TrimSpace(name) --removes any leading or trailing whitespace characters (including the newline character) from the string name.
-->fmt.Scanln(&num1) --The Scanln() function scans the input from the standard input (stdin) and stops scanning at the first newline character. It then stores the scanned input into the variable num1.
-->fmt.Println("Result: ", calculator(num1, num2, op)) --The calculator() function is called with the arguments num1, num2, and op. The result returned by the calculator() function is then printed to the standard output (stdout) using the Println() function.
-->The line reader := bufio.NewReader(os.Stdin) creates a new bufio.Reader object that reads from the standard input (os.Stdin) and assigns it to the variable reader.
    This bufio.Reader object is used later in the program to read the user's name input.
-->The line name, _ = reader.ReadString('\n') reads the user's name input from the standard input (os.Stdin) and assigns it to the variable name.
-->So the line name,_ = reader.ReadString('\n') reads a line of text from the input using the bufio.Reader object reader, assigns the resulting string to the variable name, and discards the error value. This ensures that the name variable contains only the actual name entered by the user, without any extra error values.
-->In Go, the underscore (_) is used as a "blank identifier" to discard the value of a variable that you don't need. In this case, name,_ = reader.ReadString('\n') is assigning the result of reader.ReadString('\n') to two variables: name and _. The _ variable is used to discard the second return value of reader.ReadString('\n'),
	which is an error value that we don't need in this case.

*/
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var name, op string
	var num1, num2 float32
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("Hello", name)
	fmt.Println("Enter First number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter Second number: ")
	fmt.Scanln(&num2)
	fmt.Println("Enter the operator: ")
	fmt.Scanln(&op)
	if num2 == 0 && op == "/" {
		fmt.Println("Cannot divide by zero...Enter a valid number")
		os.Exit(0)
	} else {
		fmt.Println("Result: ", calculator(num1, num2, op))
	}
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
			fmt.Println("Cannot divide by zero")
			return -9999
		}
	default:
		fmt.Println("Invalid operator")
	}
	return result
}
