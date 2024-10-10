package main

import "fmt"

const LoginToken string = "hsdfgGF9&df88" //Public

func main() {
	var username string = "Sujal"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.45454
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	//implicit type of declaring variables
	var website = "sujaldev.itch.io"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	//no var style
	numberOfUsers := 4800.0
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)

	fmt.Println(LoginToken)
}
