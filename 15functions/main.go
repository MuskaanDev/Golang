package main

import "fmt"

func main() { //entry point of program, exectues automatically
	fmt.Println("Functions")
	greet() //step 2

	result := adder(3, 5)
	greettwo()

	fmt.Println("Result is ", result)
	pro := proadder(2, 5, 8, 7)
	fmt.Println("Pro result is :", pro)

}

func adder(valueone int, valuetwo int) int { // int outside to specifiy return type value
	return valueone + valuetwo
}

func proadder(vals ...int) int { //all values of type integer
	total := 0
	for _, val := range vals {

		total += val
	}
	return total
}

//functions inside functions not allowed
//anonynous func existsts
func greettwo() {

	fmt.Println("second")
}
func greet() { //call the function
	fmt.Println("Hello") //step 1
}
