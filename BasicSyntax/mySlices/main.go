package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	fmt.Println(highScores)
	fmt.Println(len(highScores))
	// highScores[4] = 867 // Cannot be possible because range is 4 predefined

	highScores = append(highScores, 555, 666, 321) // But this way reallocation is done and all values are accomodated

	fmt.Println(highScores)
	fmt.Println(len(highScores))

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//How to remove a value from slices based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
