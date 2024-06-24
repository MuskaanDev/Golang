package main

import "fmt"

func main() {

	fmt.Println("Loops")

	days := []string{"Sun", "Mon", "Wed", "Fri", "Sat"} //slice

	fmt.Println(days)

	//for loop

	for i := 0; i < len(days); i++ { //no pre-increment operators only post
		fmt.Println(days[i])
	} //used when we know we have to traverse throught the entire list

	for i := range days { //better practice
		fmt.Println("the days are", days[i])
	}

	//for each equivalent

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	roughevalue := 1

	for roughevalue < 10 {

		if roughevalue == 2 {
			goto lco
		}
		if roughevalue == 5 {
			break //can use continue also
		}
		fmt.Println("value is:", roughevalue)
		roughevalue++
	}

	//goto statements

lco:
	fmt.Println("Jump here")
}
