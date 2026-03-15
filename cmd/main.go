package main

import (
	"log"
	"net/http"

	"kitchen-api/internal/db"
	"kitchen-api/internal/handler"
)

func main() {

	db.InitDB()

	http.HandleFunc("/ingredients", handler.IngredientHandler)
	http.HandleFunc("/produce", handler.ProduceMealHandler)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
