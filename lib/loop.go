package main

import "fmt"

func main() {
	var sum int
	sum = 0
	fmt.Println("Enter the number of elements in array")
	var n int
	fmt.Scanln(&n)
	//enter the elements of the array from user input
	num := make([]int, n)
	fmt.Printf("Enter %d integers to the array ", n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&num[i])
	}
	fmt.Println("The elements of the array are: ")
	for i := 0; i < n; i++ {
		fmt.Println(num[i])
		sum += num[i]
	}
	fmt.Println("Sum: ", sum)
	//while loop
	count := 10
	for count > -1 {
		fmt.Println("Counter---- ", count)
		//one second delay for counter
		for i := 0; i < 1000000000; i++ {
		}
		count--
	}
	fmt.Println("Shoot")
	for {
		var cmd string
		fmt.Print("Enter command: ")
		fmt.Scanln(&cmd)

		if cmd == "e" {
			break
		}
		fmt.Println("Executin the command: ", cmd)
	}
	fmt.Print("Came out of infinite loop")
}
