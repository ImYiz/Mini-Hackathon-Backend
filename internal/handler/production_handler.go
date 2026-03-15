package handler

import (
	"encoding/json"
	"net/http"

	"kitchen-api/internal/service"
)

type ProductionRequest struct {
	IngredientID int `json:"ingredient_id"`
	Quantity     int `json:"quantity"`
}

func ProduceMealHandler(w http.ResponseWriter, r *http.Request) {

	var req ProductionRequest

	json.NewDecoder(r.Body).Decode(&req)

	err := service.ProduceMeal(req.IngredientID, req.Quantity)

	if err != nil {

		http.Error(w, err.Error(), 400)

		return
	}

	w.Write([]byte("production success"))
}
