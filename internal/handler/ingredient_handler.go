package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"kitchen-api/internal/model"
	"kitchen-api/internal/repository"
)

func IngredientHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:

		var ingredient model.Ingredient

		json.NewDecoder(r.Body).Decode(&ingredient)

		repository.CreateIngredient(ingredient)

		w.WriteHeader(http.StatusCreated)

	case http.MethodGet:

		data, _ := repository.GetIngredients()

		json.NewEncoder(w).Encode(data)

	case http.MethodPut:

		idStr := r.URL.Query().Get("id")

		id, _ := strconv.Atoi(idStr)

		var ingredient model.Ingredient

		json.NewDecoder(r.Body).Decode(&ingredient)

		repository.UpdateIngredient(id, ingredient.Name, ingredient.Stock, ingredient.Unit)

		w.Write([]byte("ingredient updated"))

	case http.MethodDelete:

		idStr := r.URL.Query().Get("id")

		id, _ := strconv.Atoi(idStr)

		repository.DeleteIngredient(id)

		w.Write([]byte("ingredient deleted"))
	}
}
