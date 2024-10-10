package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sujalkbarnwal/golang/GoProjects/URLShortener/router"
)

func main() {
	fmt.Println("URL Shortener Service")
	r := router.Router()
	fmt.Println("Server is getting started...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening at port 8080 ...")
}
