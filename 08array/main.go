package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitlist [4]string

	fruitlist[0] = "apple"
	fruitlist[1] = "tomato"
	fruitlist[2] = "peas"

	fmt.Println("Fruit list is :", fruitlist)
	fmt.Println("Length of the list is:", len(fruitlist))

	var veglist = [6]string{"potato", "beans", "mushroom"}
	fmt.Println("The list is:", len(veglist))
}
