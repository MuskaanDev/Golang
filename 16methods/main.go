// alternative of classes
package main

import "fmt"

// declare a struct
type User struct { //User U capital because public works like class needs to be exported out
	Name   string //Same reason for capital here as well
	Email  string
	Status bool
	Age    int
	oneAge int //not exportable because first letter small
}

func (u User) GetStatus() {
	fmt.Println("is user active:", u.Status)
}

func (u User) NewMail() { //passes on a copy for original we need to pass the pointer
	u.Email = "test@go.dev"
	fmt.Println("Email is ", u.Email)
}

func main() {
	fmt.Println("Structs in Golang")
	//no inheritance classes in golang; no super or parent classes
	//put values in your struct
	muskaan := User{"Muskaan", "muskaan@gmail.com", true, 22, 1}
	fmt.Println(muskaan)

	//to print in an easier format,+v gives much more detail in case of structs
	fmt.Printf("muskaan's details are: %+v\n", muskaan)

	//for printing single values
	fmt.Printf("Name is %v and email is %v.", muskaan.Name, muskaan.Email)

	muskaan.GetStatus() //call the method
	muskaan.NewMail()

}

//classes are not present in golang therefore structs bnane pdte
//functions jab classes main dalo toh methods bnte hain
