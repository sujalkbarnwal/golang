package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	//type1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//type2
	for i := range days {
		fmt.Println(days[i])
	}

	//type3
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	//type4 @whileloop-type
	rogueValue := 1
	for rogueValue < 10 {

		if rogueValue == 5 {
			// break
			rogueValue++
			continue
		}

		if rogueValue == 9 {
			goto suj
		}

		fmt.Println("Value is:", rogueValue)
		rogueValue++
	}
suj:
	fmt.Println("You hit the label suj!")
}
