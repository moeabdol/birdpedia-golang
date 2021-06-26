package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moeabdol/birdpedia-golang/controllers"
	"github.com/moeabdol/birdpedia-golang/models"
	"github.com/moeabdol/birdpedia-golang/utils"
)

func main() {
	utils.LoadConfig()
	models.ConnectDatabase()

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.GetHome).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)
}
