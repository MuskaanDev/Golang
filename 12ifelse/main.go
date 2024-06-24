package main

import "fmt"

func main() {
	fmt.Println("if and else")
	var result string
	logincount := 23
	if logincount < 10 { // can't put these bracket on the next line lexer doesn't allow
		result = "Regular User"
	} else if logincount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 logins"
	}
	fmt.Println(result)

	//can create afterwards as well no need to declare variables befiore hand
	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	//used in web request handling
	//after semicolon used to assign during a web request

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}
}
