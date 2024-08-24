package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thank for the rating", input)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //You use strconv here because input is originally a string (as all data read from the terminal is in string format). To perform numerical operations on this input (like calculations or comparisons), you need to convert it to a float. strconv.ParseFloat is the tool provided by Go to handle this conversion.

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added one to your rating:", numrating+1)
	}
}
