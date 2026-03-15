package repository

import (
	"kitchen-api/internal/db"
	"kitchen-api/internal/model"
)

func CreateIngredient(i model.Ingredient) error {

	query := `
	INSERT INTO ingredients(name,stock,unit)
	VALUES($1,$2,$3)
	`

	_, err := db.DB.Exec(query, i.Name, i.Stock, i.Unit)

	return err
}

func GetIngredients() ([]model.Ingredient, error) {

	rows, err := db.DB.Query("SELECT id,name,stock,unit FROM ingredients")

	if err != nil {
		return nil, err
	}

	var ingredients []model.Ingredient

	for rows.Next() {

		var i model.Ingredient

		rows.Scan(&i.ID, &i.Name, &i.Stock, &i.Unit)

		ingredients = append(ingredients, i)
	}

	return ingredients, nil
}

func UpdateStock(id int, stock int) error {

	_, err := db.DB.Exec(
		"UPDATE ingredients SET stock=$1 WHERE id=$2",
		stock,
		id,
	)

	return err
}

func GetIngredientByID(id int) (model.Ingredient, error) {

	var i model.Ingredient

	err := db.DB.QueryRow(
		"SELECT id,name,stock,unit FROM ingredients WHERE id=$1",
		id,
	).Scan(&i.ID, &i.Name, &i.Stock, &i.Unit)

	return i, err
}

func UpdateIngredient(id int, name string, stock int, unit string) error {

	query := `
	UPDATE ingredients
	SET name=$1, stock=$2, unit=$3
	WHERE id=$4
	`

	_, err := db.DB.Exec(query, name, stock, unit, id)

	return err
}

func DeleteIngredient(id int) error {

	query := `DELETE FROM ingredients WHERE id=$1`

	_, err := db.DB.Exec(query, id)

	return err
}
