package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is", fruitList)
	fmt.Println("Length of fruit list is", len(fruitList))

	var vegList = [5]string{"Potato", "Cauliflower", "Okra"}
	fmt.Println("Vegy list is", vegList)
	fmt.Println("Length of vegy list is", len(vegList))
}
