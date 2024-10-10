package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(5, 3)
	fmt.Println("result is:", result)

	proRes, myMsg := proAdder(3, 2, 1, 4)
	fmt.Println("Pro result is:", proRes, myMsg)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) { //variadic function
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "This function is proAdder ;)"
}

func greeter() {
	fmt.Println("Welcome user!")
}
