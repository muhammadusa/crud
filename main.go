package main

import (
	// "fmt"
	"app/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/products", controllers.GetProduct).Methods("GET")
	r.HandleFunc("/products", controllers.PostProduct).Methods("POST")
	r.HandleFunc("/products", controllers.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/products", controllers.PutProduct).Methods("PUT")

	http.ListenAndServe(":8080", r)

}
