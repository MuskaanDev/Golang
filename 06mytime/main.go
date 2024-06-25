package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	create := time.Date(2023, time.December, 24, 23, 23, 0, 0, time.Local)

	fmt.Println(create.Format("01-02-2006 Monday"))

}
