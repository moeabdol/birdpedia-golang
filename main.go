package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moeabdol/birdpedia-golang/controllers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.GetHome).Methods(http.MethodGet)
	http.ListenAndServe(":8080", r)
}
