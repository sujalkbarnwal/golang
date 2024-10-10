package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	//comma ok //comma error
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:", input)
	fmt.Printf("Rating is of type: %T \n", input)

}
