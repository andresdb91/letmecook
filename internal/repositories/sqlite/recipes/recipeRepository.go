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
	next := func() (*QueryResult[DBRecipe], error) {
		return GetQueryPage[DBRecipe](rr.db, "SELECT * FROM recipes", recipeItem, recipeRows, pageSize, 0)
	}
	parser := func(dbRecipePage *[]*DBRecipe) ([]*recipes.Recipe, error) {
		return ParseDBRecipeList(*dbRecipePage)
	}
	result, err := ParseQueryPage[DBRecipe, recipes.Recipe](next, parser)
	if err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

// Recibe un `next` y un parser; devuelve un `PagedList[T]` parseado para business
func ParseQueryPage[F any, T any](query func() (*QueryResult[F], error), parser func(dbPage *[]*F) ([]*T, error)) (*commons.PagedList[T], error) {
	page, queryErr := query()
	if queryErr != nil {
		return nil, queryErr
	}
	items, parserErr := parser(&page.items)
	if parserErr != nil {
		return nil, parserErr
	}
	pagedResult := commons.PagedList[T]{
		Items:       items,
		PageSize:    page.pageSize,
		CurrentPage: page.currentPage,
		HasNext:     page.next != nil,
		GetNextPage: func() (*commons.PagedList[T], error) {
			return ParseQueryPage[F, T](page.next, parser)
		},
	}
	return &pagedResult, nil
}

func ParseDBRecipeList(dbRecipeList []*DBRecipe) ([]*recipes.Recipe, error) {
	var recipeList []*recipes.Recipe
	for _, dbRecipe := range dbRecipeList {
		recipeList = append(recipeList, ParseRecipe(dbRecipe))
	}
	return recipeList, nil
}

func ParseRecipe(dbRecipe *DBRecipe) *recipes.Recipe {
	return &recipes.Recipe{
		ID:                 uuid.MustParse(dbRecipe.ID),
		Name:               dbRecipe.Name,
		ResultingComponent: ParseComponent(dbRecipe.ResultingComponentID),
	}
}

func ParseComponent(dbComponentID string) *recipes.Component {
	return nil
}

type QueryResult[T any] struct {
	items       []*T
	next        func() (*QueryResult[T], error)
	currentPage int
	pageSize    int
}

func GetQueryPage[T any](db *sql.DB, query string, scanItem *T, scanRows []any, pageSize int, pageNumber int) (*QueryResult[T], error) {
	query = query + " LIMIT ? OFFSET ?"
	rows, err := db.Query(query, pageSize+1, pageNumber*pageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resultList []*T
	// var pagedRecipes commons.PagedList[T]
	var scanError error
	for rows.Next() {
		if scanError = rows.Scan(scanRows...); scanError != nil {
			// return &pagedRecipes, err
			break
		}
		resultList = append(resultList, scanItem)
	}

	var next func() (*QueryResult[T], error)
	if len(resultList) > pageSize {
		next = func() (*QueryResult[T], error) {
			return GetQueryPage[T](db, query, scanItem, scanRows, pageSize, pageNumber+1)
		}
	} else {
		next = nil
	}

	result := QueryResult[T]{
		items:       resultList,
		next:        next,
		currentPage: pageNumber,
		pageSize:    pageSize,
	}

	if scanError != nil {
		return &result, scanError
	} else {
		return &result, nil
	}
}
