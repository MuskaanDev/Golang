package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("files")
	content := "this need to go in a file"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err) //shuts down the process of the program and shows u the error u are facing
	}

	length, err := io.WriteString(file, content)
}
