package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string) //maps, bracket has key,next is value

	languages["JS"] = "Java"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages) //prints languages using key
	//prints a particular language using key
	fmt.Println("JS shorts for :", languages["JS"])

	delete(languages, "RB") //delete using key
	fmt.Println("List of all languages:", languages)

	for key, value := range languages {
		fmt.Printf("For key %v, values is %v\n", key, value)
	} //for loop equivalent of for each loop so for this key this value range is languages
	//print

	for _, value := range languages {
		fmt.Printf("For key v, values is %v\n", value)
	} //can ignore any value using error, ok syntax

}
