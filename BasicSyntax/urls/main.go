package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=golang&paymentid=jdsfHFJK7875FG"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery) //params or parameters

	qparams := result.Query()
	fmt.Printf("Query params type is %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/learncss",
		RawPath: "user=sujal",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
