package recipes

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"

	recipes "github.com/andresdb91/letmecook/internal/business/recipes"
	"github.com/andresdb91/letmecook/pkg/commons"
)

type SQLiteRecipeRepository struct {
	db *sql.DB
}

func (rr *SQLiteRecipeRepository) GetRecipeByID(id string) (*recipes.Recipe, error) {
	return nil, nil
}

func (rr *SQLiteRecipeRepository) GetRecipesByComponents(components []string) ([]*recipes.Recipe, error) {
	return nil, nil
}

func (rr *SQLiteRecipeRepository) GetRecipesByKeywords(keywords []string) ([]*recipes.Recipe, error) {
	return nil, nil
}

func (rr *SQLiteRecipeRepository) GetRecipesByName(name string) ([]*recipes.Recipe, error) {
	return nil, nil
}

func (rr *SQLiteRecipeRepository) GetRecipesByTags(tags []string) ([]*recipes.Recipe, error) {
	return nil, nil
}

func (rr *SQLiteRecipeRepository) GetTopRatedRecipes() ([]*recipes.Recipe, error) {
	return nil, nil
}

func (rr *SQLiteRecipeRepository) GetTopSearchedRecipes() ([]*recipes.Recipe, error) {
	return nil, nil
}

func (rr *SQLiteRecipeRepository) GetNewestRecipes() ([]*recipes.Recipe, error) {
	return nil, nil
}

func (rr *SQLiteRecipeRepository) GetRandomRecipe() (*recipes.Recipe, error) {
	return nil, nil
}

func (rr *SQLiteRecipeRepository) GetAllRecipes() (*commons.PagedList[recipes.Recipe], error) {
	pageSize := 10
	recipeItem := new(DBRecipe)
	recipeRows := []any{&recipeItem.ID, &recipeItem.Name, &recipeItem.ResultingComponentID}

	next := func() (*commons.QueryResult[DBRecipe], error) {
		return GetQueryPage[DBRecipe](rr.db, "SELECT * FROM recipes", recipeItem, recipeRows, pageSize, 0)
	}
	parser := func(dbRecipePage *[]*DBRecipe) ([]*recipes.Recipe, error) {
		return parseDBRecipeList(*dbRecipePage)
	}

	result, err := commons.ParseQueryPage[DBRecipe, recipes.Recipe](next, parser)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

func parseDBRecipeList(dbRecipeList []*DBRecipe) ([]*recipes.Recipe, error) {
	var recipeList []*recipes.Recipe
	for _, dbRecipe := range dbRecipeList {
		parsedRecipe, err := parseRecipe(dbRecipe)
		if err != nil {
			return nil, err
		}
		recipeList = append(recipeList, parsedRecipe)
	}
	return recipeList, nil
}

func parseRecipe(dbRecipe *DBRecipe) (*recipes.Recipe, error) {
	component, err := ComponentRepository.GetComponentByID(dbRecipe.ResultingComponentID)
	ingredients, err := IngredientRepository.GetIngredientsByRecipeID(dbRecipe.ID)
	steps, err := StepRepository.GetStepsByRecipeID(dbRecipe.ID)
	tags, err := TagRepository.GetTagsByOwnerID(dbRecipe.ID)
	if err != nil {
		return nil, err
	}

	return &recipes.Recipe{
		ID:                 uuid.MustParse(dbRecipe.ID),
		Name:               dbRecipe.Name,
		ResultingComponent: component,
		Ingredients:        ingredients,
		Steps:              steps,
		Tags:               tags,
	}, nil
}
