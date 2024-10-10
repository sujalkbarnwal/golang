package main

import "fmt"

//defer works as LIFO

func main() {
	defer fmt.Println("World") //First-In
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer() //Last-In
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
