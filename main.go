package main

import (
	"log"
	"movie-recommender/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/title", handler.Suggest).Methods("POST")

	log.Println("Starting the application...")
	log.Println("Service active and listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
