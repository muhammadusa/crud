package controllers

import (
	"app/database"
	"encoding/json"
	"net/http"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(database.GetProducts())
}
