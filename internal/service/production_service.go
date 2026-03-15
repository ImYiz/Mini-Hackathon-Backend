package service

import (
	"errors"

	"kitchen-api/internal/repository"
)

func ProduceMeal(ingredientID int, qty int) error {

	ingredient, err := repository.GetIngredientByID(ingredientID)

	if err != nil {
		return err
	}

	if ingredient.Stock < qty {
		return errors.New("stock not enough")
	}

	newStock := ingredient.Stock - qty

	return repository.UpdateStock(ingredientID, newStock)
}
