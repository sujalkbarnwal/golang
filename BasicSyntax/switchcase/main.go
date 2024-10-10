package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch-case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice value:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Move 2 spots")
	case 3:
		fmt.Println("Move 3 spots")
	case 4:
		fmt.Println("Move 4 spots")
		fallthrough
	case 5:
		fmt.Println("Move 5 spots")
	case 6:
		fmt.Println("Move 6 spots and you can roll the dice again")
		fallthrough
	default:
		fmt.Println("What was that!")
	}
}
