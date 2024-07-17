package main

import "fmt"

func main() {
	defer fmt.Println("world") //even if u write at the top it will execute at the last just before the last curly braces
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")
	//output hello,two,one,world LIFO
	mydefer() //after hello this wull go next because yaha pe defer nahi hai

	//hello,43210.two.one.world
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
