package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev" //declare it in public so that the url is easy to modify
func main() {
	fmt.Println("web requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response) //type will be *http.Response, this pointer indicates we are getting original response and not a copy of it

	defer response.Body.Close() //caller's responsibiliy=ty to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes) //The response body is initially read as a slice of bytes ([]byte). This is because the response could contain any type of data, such as binary data, images, or plain text. However, if you're dealing with text data (like HTML, JSON, or plain text), it's more convenient to work with it as a string.
	fmt.Println(content)
}
