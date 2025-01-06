package recipes

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"

	recipes "github.com/andresdb91/letmecook/internal/business/recipes"
	"github.com/andresdb91/letmecook/pkg/commons"
)

type SQLiteStepRepository struct {
	db *sql.DB
}

func (r SQLiteStepRepository) GetStepsByRecipeID(recipeID string) (*commons.PagedList[recipes.Step], error) {
	pageSize := 10
	stepItem := new(DBStep)
	stepRows := []any{&stepItem.ID, &stepItem.Description, &stepItem.Order}

	next := func() (*commons.QueryResult[DBStep], error) {
		return GetQueryPage[DBStep](r.db, "SELECT * FROM steps WHERE recipe_id = ?", stepItem, stepRows, pageSize, recipeID)
	}

	return GetPagedSteps(next)
}
