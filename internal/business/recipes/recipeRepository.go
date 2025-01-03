package recipes

import "github.com/andresdb91/letmecook/pkg/commons"

type ComponentRepository interface {
	GetComponentByID(id string) (*Component, error)
	GetComponentsByTags(tags []string) ([]*Component, error)
}

type UnitRepository interface {
}

type QuantityRepository interface {
}

type IngredientRepository interface {
}

type StepRepository interface {
}

type RecipeRepository interface {
	GetRecipeByID(id string) (*Recipe, error)
	GetRecipesByComponents(components []string) ([]*Recipe, error)
	GetRecipesByKeywords(keywords []string) ([]*Recipe, error)
	GetRecipesByName(name string) ([]*Recipe, error)
	GetRecipesByTags(tags []string) ([]*Recipe, error)
	GetTopRatedRecipes() ([]*Recipe, error)
	GetTopSearchedRecipes() ([]*Recipe, error)
	GetNewestRecipes() ([]*Recipe, error)
	GetRandomRecipe() (*Recipe, error)
	GetAllRecipes() (*commons.PagedList[Recipe], error)
}
