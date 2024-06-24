// alternative of classes
package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	//no inheritance classes in golang; no super or parent classes
	muskaan := User{"Muskaan", "muskaan@gmail.com", true, 22}
	fmt.Println(muskaan)

	//to print in an easier format,+v gives much more detail in case of structs
	fmt.Printf("muskaan's details are: %+v\n", muskaan)

	//for printing single values
	fmt.Printf("Name is %v and email is %v.", muskaan.Name, muskaan.Email)

}

type User struct { //User U capital because public works like class needs to be exported out
	Name   string //Same reason for capital here as well
	Email  string
	Status bool
	Age    int
}
