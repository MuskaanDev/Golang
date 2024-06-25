package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int

	fmt.Println("The value of pinter is", ptr)

	var a = 23
	ptr = &a
	*ptr = *ptr * 2
	fmt.Println("The value of the pointer is", *ptr)
}
