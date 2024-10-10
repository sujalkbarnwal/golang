package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://example.com/"

func main() {
	fmt.Println("Example.com web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of response: %T", response)

	defer response.Body.Close() //Highly important to close the connection

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
