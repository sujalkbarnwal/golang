package router

import (
	"github.com/gorilla/mux"
	"github.com/sujalkbarnwal/golang/GoProjects/URLShortener/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.GetHomePage).Methods("GET")
	router.HandleFunc("/shorten", controller.CreateShortenUrl).Methods("POST")
	router.HandleFunc("/{shortcode}", controller.ReturnOriginalUrl).Methods("GET")

	return router
}
