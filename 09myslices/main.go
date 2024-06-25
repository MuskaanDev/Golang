package main

import "fmt"

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitlist = []string{"apple", "tomato", "peas"}
	fmt.Printf("Type of fruitlist is %T\n", fruitlist)

	fruitlist = append(fruitlist, "mango", "banana")
	fmt.Println(fruitlist)

	fruitlist = fruitlist[1:3] // Corrected slicing
	fmt.Println(fruitlist)

	highscore := make([]int, 4)
	highscore[0] = 3
	highscore[1] = 4
	fmt.Println(highscore) // Added this line to print highscore

	var courses = []string{"react", "java", "golang", "cpp"}
	fmt.Println(courses)

	var index int = 2
	// Corrected slice creation and appending
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
