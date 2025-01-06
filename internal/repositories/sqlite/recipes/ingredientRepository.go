package recipes

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"

	recipes "github.com/andresdb91/letmecook/internal/business/recipes"
	"github.com/andresdb91/letmecook/pkg/commons"
)

type SQLiteIngredientRepository struct {
	db *sql.DB
}

func (r SQLiteIngredientRepository) GetIngredientsByRecipeID(recipeID string) (*commons.PagedList[recipes.Ingredient], error) {
	pageSize := 10
	ingredientItem := new(DBIngredient)
	ingredientRows := []any{&ingredientItem.ID, &ingredientItem.Name, &ingredientItem.Quantity, &ingredientItem.UnitID}

	next := func() (*commons.QueryResult[DBIngredient], error) {
		return GetQueryPage[DBIngredient](r.db, "SELECT * FROM ingredients WHERE recipe_id = ?", ingredientItem, ingredientRows, pageSize, recipeID)
	}

	return GetPagedIngredients(next)
}
